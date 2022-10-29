package radio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net/url"
	"strconv"
	"unicode"

	"github.com/gen2brain/malgo"
	"matsu.dev/matsuri-contest-log/hamlib"
	"matsu.dev/matsuri-contest-log/morsecode"
)

type RadioMode uint32

const (
	RadioMode_AM      RadioMode = 1 << 0
	RadioMode_CW      RadioMode = 1 << 1
	RadioMode_USB     RadioMode = 1 << 2
	RadioMode_LSB     RadioMode = 1 << 3
	RadioMode_RTTY    RadioMode = 1 << 4
	RadioMode_FM      RadioMode = 1 << 5
	RadioMode_WFM     RadioMode = 1 << 6
	RadioMode_CWR     RadioMode = 1 << 7
	RadioMode_RTTYR   RadioMode = 1 << 8
	RadioMode_AMS     RadioMode = 1 << 9
	RadioMode_PKTLSB  RadioMode = 1 << 10
	RadioMode_PKTUSB  RadioMode = 1 << 11
	RadioMode_PKTFM   RadioMode = 1 << 12
	RadioMode_ECSSUSB RadioMode = 1 << 13
	RadioMode_ECSSLSB RadioMode = 1 << 14
	RadioMode_FAX     RadioMode = 1 << 15
	RadioMode_SAM     RadioMode = 1 << 16
	RadioMode_SAL     RadioMode = 1 << 17
	RadioMode_SAH     RadioMode = 1 << 18
	RadioMode_DSB     RadioMode = 1 << 19
	RadioMode_FMN     RadioMode = 1 << 21
)

type Radio struct {
	rxFreq int64
	rxMode RadioMode
	txFreq int64
	txMode RadioMode

	rig     *hamlib.Rig
	audioTx *malgo.Device

	txAudioBuffer bytes.Buffer
	txSampleRate  uint32
}

func (r *Radio) InitAudioOutput(ctx malgo.Context, device string) error {
	devices, err := ctx.Devices(malgo.Playback)
	if err != nil {
		return fmt.Errorf("failed to get devices: %v", err)
	}

	for _, dev := range devices {
		if dev.ID.String() == device {
			r.txSampleRate = dev.MaxSampleRate
			r.audioTx, err = malgo.InitDevice(ctx, malgo.DeviceConfig{
				DeviceType:         malgo.Playback,
				SampleRate:         dev.MaxSampleRate,
				PeriodSizeInFrames: dev.MaxSampleRate / 100,
				Resampling: malgo.ResampleConfig{
					Algorithm: malgo.ResampleAlgorithmSpeex,
					Speex: malgo.ResampleSpeexConfig{
						Quality: 5,
					},
				},
				Playback: malgo.SubConfig{
					DeviceID:  dev.ID.Pointer(),
					Format:    malgo.FormatS16,
					Channels:  1,
					ShareMode: malgo.Shared,
				},
			},
				malgo.DeviceCallbacks{
					Data: r.audioSampleCallbackfunc,
				})
			if err != nil {
				return fmt.Errorf("failed to open device for sending audio: %v", err)
			}
			break
		}
	}

	return nil
}

func (r *Radio) audioSampleCallbackfunc(pOutputSample, pInputSamples []byte, framecount uint32) {
	buflen := len(pOutputSample)
	n, _ := r.txAudioBuffer.Read(pOutputSample)
	if n != buflen {
		for i := n; i < buflen; i++ {
			pOutputSample[i] = 0
		}
		go r.audioTx.Stop()
	}
}

func (r *Radio) Open(uristr string) error {
	path, err := url.Parse(uristr)
	if err != nil {
		return fmt.Errorf("failed to open radio: %v", err)
	}
	model := path.Query().Get("model")

	if r.rig != nil {
		r.rig.Close()
		r.rig.Cleanup()
		r.rig = nil
	}

	{
		models := hamlib.ListModels()
		for _, m := range models {
			if model == fmt.Sprintf("%s %s", m.Manufacturer, m.Manufacturer) {
				r.rig = &hamlib.Rig{}
				err := r.rig.Init(m.ModelID)
				if err != nil {
					return fmt.Errorf("failed to open radio: %v", err)
				}
				break
			}
		}

		if r.rig == nil {
			return fmt.Errorf("failed to open radio: unsupported radio")
		}
	}

	{
		rigPort := hamlib.Port{}
		switch path.Scheme {
		case "serial":
			rigPort.RigPortType = hamlib.RigPortSerial
			if path.Query().Has("baudrate") {
				rigPort.Baudrate, _ = strconv.Atoi(path.Query().Get("baudrate"))
			}
			if path.Query().Has("databits") {
				rigPort.Databits, _ = strconv.Atoi(path.Query().Get("databits"))
			}
			if path.Query().Has("stopbits") {
				rigPort.Stopbits, _ = strconv.Atoi(path.Query().Get("stopbits"))
			}
			if path.Query().Has("parity") {
				switch path.Query().Get("parity") {
				case "none":
					rigPort.Parity = hamlib.ParityNone
				case "even":
					rigPort.Parity = hamlib.ParityEven
				case "odd":
					rigPort.Parity = hamlib.ParityOdd
				}
			}
			if path.Query().Has("handshake") {
				switch path.Query().Get("handshake") {
				case "none":
					rigPort.Handshake = hamlib.HandshakeNone
				case "rtscts":
					rigPort.Handshake = hamlib.HandshakeRTSCTS
				}
			}
		default:
			return fmt.Errorf("failed to open radio: unsupported port")
		}
		r.rig.SetPort(rigPort)
	}

	if err := r.rig.Open(); err != nil {
		r.rig.Close()
		r.rig.Cleanup()
		r.rig = nil
		return err
	}

	return nil
}

func (r *Radio) Close() error {
	r.rig.Close()
	r.rig.Cleanup()
	r.rig = nil
	return nil
}

func (r *Radio) UpdateFreqMode() error {
	if r.rig == nil {
		return nil
	}
	if r.rxFreq != 0 {
		r.rig.SetMode(hamlib.VFOCurrent, hamlib.Mode(r.rxMode), -1)
		r.rig.SetFreq(hamlib.VFOCurrent, float64(r.rxFreq))
	}
	if r.txFreq != 0 {
		if r.txMode == RadioMode_CW {
			r.rig.SetSplitMode(hamlib.VFOCurrent, hamlib.ModeLSB, -1)
			r.rig.SetSplitFreq(hamlib.VFOCurrent, float64(r.rxFreq+1500))
		} else {
			r.rig.SetSplitMode(hamlib.VFOCurrent, hamlib.Mode(r.rxMode), -1)
			r.rig.SetSplitFreq(hamlib.VFOCurrent, float64(r.rxFreq))
		}
	}
	return nil
}

func writeSilence(wr io.ByteWriter, frame uint32) {
	for i := uint32(0); i < frame*2; i++ {
		wr.WriteByte(0)
	}
}

func writeWave(wr io.Writer, maxVal uint32, sr uint32, length uint32, rising uint32) {
	for i := uint32(0); i < length; i++ {
		ti := float64(i) / float64(sr)
		if i < rising {
			val := int16((1 - (math.Cos(math.Pi*float64(i)/float64(rising))+1)*0.5) * math.Cos(math.Pi*ti/1500.) * float64(maxVal))
			binary.Write(wr, binary.LittleEndian, val)
		} else {
			val := int16(math.Cos(math.Pi*ti/1500.) * float64(maxVal))
			binary.Write(wr, binary.LittleEndian, val)
		}
	}
	for i := uint32(0); i < rising; i++ {
		ti := float64(i+length) / float64(sr)
		val := int16(((math.Cos(math.Pi*float64(i)/float64(rising)) + 1) * 0.5) * math.Cos(math.Pi*ti/1500.) * float64(maxVal))
		binary.Write(wr, binary.LittleEndian, val)
	}
}

func (r *Radio) SendCW(chars []byte, basewpn uint32) {
	r.txAudioBuffer.Reset()

	wpm := basewpn
	for _, ch := range chars {
		if ch == 0xfe {
			wpm += 2
			continue
		} else if ch == 0xff {
			wpm -= 2
			continue
		}
		code := morsecode.GetCode(unicode.ToLower(rune(ch)))
		if len(code) != 0 {
			for _, c := range code {
				if c > 0 {
					writeWave(&r.txAudioBuffer, 32000, r.txSampleRate, uint32(float64(r.txSampleRate*uint32(c))*1.2/float64(wpm)), r.txSampleRate/100)
					writeSilence(&r.txAudioBuffer, r.txSampleRate/100-uint32(float64(r.txSampleRate)*1.2/float64(wpm)))
				} else {
					writeSilence(&r.txAudioBuffer, uint32(float64(r.txSampleRate*uint32(-c))*1.2/float64(wpm)))
				}
			}
			writeSilence(&r.txAudioBuffer, uint32(float64(r.txSampleRate*3)*1.2/float64(wpm)))
		}
	}

	if r.audioTx != nil && !r.audioTx.IsStarted() {
		r.audioTx.Start()
	}
}
