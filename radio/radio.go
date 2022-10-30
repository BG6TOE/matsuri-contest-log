package radio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"net/url"
	"strconv"
	"sync"
	"time"
	"unicode"

	"github.com/gen2brain/malgo"
	"github.com/sirupsen/logrus"
	"matsu.dev/matsuri-contest-log/hamlib"
	"matsu.dev/matsuri-contest-log/morsecode"
	pb "matsu.dev/matsuri-contest-log/proto"
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

func (m RadioMode) ToProtoMode() pb.RadioMode {
	switch m {
	case RadioMode_AM:
		return pb.RadioMode_AM
	case RadioMode_CW:
		return pb.RadioMode_CW
	case RadioMode_USB:
		return pb.RadioMode_CWR
	case RadioMode_LSB:
		return pb.RadioMode_LSB
	case RadioMode_RTTY:
		return pb.RadioMode_DATAU
	case RadioMode_FM:
		return pb.RadioMode_FM
	case RadioMode_WFM:
		return pb.RadioMode_FM
	case RadioMode_CWR:
		return pb.RadioMode_CWR
	case RadioMode_RTTYR:
		return pb.RadioMode_DATAL
	case RadioMode_AMS:
		return pb.RadioMode_AM
	case RadioMode_PKTLSB:
		return pb.RadioMode_DATAL
	case RadioMode_PKTUSB:
		return pb.RadioMode_DATAU
	case RadioMode_PKTFM:
		return pb.RadioMode_FM
	case RadioMode_ECSSUSB:
		return pb.RadioMode_DATAU
	case RadioMode_ECSSLSB:
		return pb.RadioMode_DATAL
	case RadioMode_FAX:
		return pb.RadioMode_DATAU
	case RadioMode_SAM:
		return pb.RadioMode_UNKNOWN
	case RadioMode_SAL:
		return pb.RadioMode_UNKNOWN
	case RadioMode_SAH:
		return pb.RadioMode_UNKNOWN
	case RadioMode_DSB:
		return pb.RadioMode_UNKNOWN
	case RadioMode_FMN:
		return pb.RadioMode_FM
	default:
		return pb.RadioMode_UNKNOWN
	}
}

func ToRadioMode(m pb.RadioMode) RadioMode {
	switch m {
	case pb.RadioMode_UNKNOWN:
		return 0
	case pb.RadioMode_CW:
		return RadioMode_CW
	case pb.RadioMode_CWR:
		return RadioMode_CWR
	case pb.RadioMode_LSB:
		return RadioMode_LSB
	case pb.RadioMode_USB:
		return RadioMode_USB
	case pb.RadioMode_AM:
		return RadioMode_AM
	case pb.RadioMode_FM:
		return RadioMode_FM
	case pb.RadioMode_DATAL:
		return RadioMode_PKTLSB
	case pb.RadioMode_DATAU:
		return RadioMode_PKTUSB
	default:
		return 0
	}
}

type RadioSetting struct {
	Backend string `json:"backend"`
	Model   string `json:"model"`
	Uri     string `json:"uri"`
}

type Radio struct {
	RadioSetting

	rxFreq    int64
	rxMode    RadioMode
	txFreq    int64
	txMode    RadioMode
	isTx      bool
	isLocalTx bool

	rig     *hamlib.Rig
	audioTx *malgo.Device

	txAudioBuffer bytes.Buffer
	txSampleRate  uint32

	lock        sync.Mutex
	refreshChan chan struct{}
}

func NewRadio(s RadioSetting) *Radio {
	return &Radio{
		RadioSetting: s,
	}
}

func (r *Radio) InitAudioOutput(ctx malgo.Context, device string) error {
	devices, err := ctx.Devices(malgo.Playback)
	if err != nil {
		return fmt.Errorf("failed to get devices: %v", err)
	}

	for _, dev := range devices {
		if dev.ID.String() == device {
			r.txSampleRate = 48000
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

	if r.audioTx == nil {
		return fmt.Errorf("failed to open device for sending audio: device not found")
	}

	return nil
}

func (r *Radio) stopTx() {
	r.lock.Lock()
	r.audioTx.Stop()
	r.rig.SetPtt(hamlib.VFOCurrent, 0)
	if r.txMode == RadioMode_CW {
		r.rig.SetMode(hamlib.VFOCurrent, hamlib.ModeCW, -1)
		r.rig.SetFreq(hamlib.VFOCurrent, float64(r.txFreq))
	}
	r.isLocalTx = false
	r.lock.Unlock()
}

func (r *Radio) audioSampleCallbackfunc(pOutputSample, pInputSamples []byte, framecount uint32) {
	buflen := len(pOutputSample)
	n, _ := r.txAudioBuffer.Read(pOutputSample)
	if n != buflen {
		for i := n; i < buflen; i++ {
			pOutputSample[i] = 0
		}
		go r.stopTx()
	}
}

func (r *Radio) Open() error {
	path, err := url.Parse(r.Uri)
	if err != nil {
		return fmt.Errorf("failed to open radio: %v", err)
	}

	logrus.Infof("Opening %s at %s", r.Model, path.Path)

	r.Close()

	r.refreshChan = make(chan struct{})

	var tmprig *hamlib.Rig

	{
		models := hamlib.ListModels()
		for _, m := range models {
			if r.Model == fmt.Sprintf("%s %s", m.Manufacturer, m.Model) {
				tmprig = &hamlib.Rig{}
				err := tmprig.Init(m.ModelID)
				if err != nil {
					return fmt.Errorf("failed to init radio: %v", err)
				}
				break
			}
		}

		if tmprig == nil {
			return fmt.Errorf("failed to init radio: unsupported radio %s", r.Model)
		}
	}

	defer func() {
		if tmprig != nil {
			tmprig.Cleanup()
		}
	}()

	{
		rigPort := hamlib.Port{}
		switch path.Scheme {
		case "serial":
			rigPort.Portname = path.Path
			rigPort.RigPortType = hamlib.RigPortSerial
			if path.Query().Has("baudrate") {
				rigPort.Baudrate, _ = strconv.Atoi(path.Query().Get("baudrate"))
			} else {
				rigPort.Baudrate = 19200
			}
			if path.Query().Has("databits") {
				rigPort.Databits, _ = strconv.Atoi(path.Query().Get("databits"))
			} else {
				rigPort.Databits = tmprig.Caps.SerialDataBits
			}
			if path.Query().Has("stopbits") {
				rigPort.Stopbits, _ = strconv.Atoi(path.Query().Get("stopbits"))
			} else {
				rigPort.Stopbits = tmprig.Caps.SerialStopBits
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
			} else {
				rigPort.Parity = hamlib.Parity(tmprig.Caps.SerialParity)
			}
			if path.Query().Has("handshake") {
				switch path.Query().Get("handshake") {
				case "none":
					rigPort.Handshake = hamlib.HandshakeNone
				case "rtscts":
					rigPort.Handshake = hamlib.HandshakeRTSCTS
				}
			} else {
				rigPort.Handshake = hamlib.Handshake(tmprig.Caps.SerialHandshake)
			}
		case "tcp":
			rigPort.Portname = path.Host
			rigPort.RigPortType = hamlib.RigPortNetwork
		default:
			return fmt.Errorf("failed to open radio: unsupported port")
		}
		if err := tmprig.SetPort(rigPort); err != nil {
			return fmt.Errorf("failed to set port: %v", err)
		}
	}

	if err := tmprig.Open(); err != nil {
		return err
	}

	r.rig = tmprig
	tmprig = nil

	go func(closeChan chan struct{}) {
		for {
			select {
			case <-time.After(3 * time.Second):
				func() {
					r.lock.Lock()
					defer r.lock.Unlock()
					if r.isLocalTx {
						r.isTx = r.isLocalTx
						return
					}

					isTx, _ := r.rig.GetPtt(hamlib.VFOCurrent)
					freq, _ := r.rig.GetFreq(hamlib.VFOCurrent)
					mode, _, _ := r.rig.GetMode(hamlib.VFOCurrent)

					logrus.Infof("VFO %v %v", int64(freq), hamlib.ModeName[mode])

					r.isTx = isTx != 0
					r.rxFreq = int64(freq)
					r.rxMode = RadioMode(mode)
				}()
			case <-closeChan:
				return
			}
		}
	}(r.refreshChan)

	return nil
}

func (r *Radio) Close() error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.refreshChan != nil {
		close(r.refreshChan)
	}
	if r.rig != nil {
		r.rig.Close()
		r.rig.Cleanup()
		r.rig = nil
	}
	return nil
}

func (r *Radio) UpdateFreqMode() error {
	if r.rig == nil {
		return nil
	}

	r.lock.Lock()
	defer r.lock.Unlock()

	r.rig.SetVfo(hamlib.VFOA)
	if r.rxFreq != 0 {
		r.rig.SetMode(hamlib.VFOCurrent, hamlib.Mode(r.rxMode), -1)
		r.rig.SetFreq(hamlib.VFOCurrent, float64(r.rxFreq))
	}
	return nil
}

func writeSilence(wr io.Writer, frame uint32) {
	for i := uint32(0); i < frame*2; i++ {
		wr.Write([]byte{0})
	}
}

func writeWave(wr io.Writer, maxVal uint32, sr uint32, length uint32, rising uint32) {
	for i := uint32(0); i < length; i++ {
		ti := float64(i) / float64(sr)
		wave := math.Cos(math.Pi * ti * 3000.)
		if i < rising {
			val := int16((1 - (math.Cos(math.Pi*float64(i)/float64(rising))+1)*0.5) * wave * float64(maxVal))
			binary.Write(wr, binary.LittleEndian, val)
		} else {
			val := int16(wave * float64(maxVal))
			binary.Write(wr, binary.LittleEndian, val)
		}
	}
	for i := uint32(0); i < rising; i++ {
		ti := float64(i+length) / float64(sr)
		wave := math.Cos(math.Pi * ti * 3000.)
		val := int16(((math.Cos(math.Pi*float64(i)/float64(rising)) + 1) * 0.5) * wave * float64(maxVal))
		binary.Write(wr, binary.LittleEndian, val)
	}
}

func (r *Radio) SendCW(chars []byte, basewpn uint32) {
	r.txAudioBuffer.Reset()

	wpm := basewpn
	writeSilence(&r.txAudioBuffer, uint32(float64(r.txSampleRate*3)*1.2/float64(wpm)))
	for _, ch := range chars {
		if ch == '{' {
			wpm += 2
			continue
		} else if ch == '}' {
			wpm -= 2
			continue
		}
		code := morsecode.GetCode(unicode.ToLower(rune(ch)))
		if len(code) != 0 {
			for _, c := range code {
				if c > 0 {
					writeWave(&r.txAudioBuffer, 32000, r.txSampleRate, uint32(float64(r.txSampleRate*uint32(c))*1.2/float64(wpm)), r.txSampleRate/100)
					writeSilence(&r.txAudioBuffer, uint32(float64(r.txSampleRate)*1.2/float64(wpm))-r.txSampleRate/100)
				} else {
					writeSilence(&r.txAudioBuffer, uint32(float64(r.txSampleRate)*float64(-c)*1.2/float64(wpm)))
				}
			}
			writeSilence(&r.txAudioBuffer, uint32(float64(r.txSampleRate*3)*1.2/float64(wpm)))
		}
	}

	if r.audioTx != nil && !r.audioTx.IsStarted() {
		r.lock.Lock()
		r.isLocalTx = true
		if r.txMode == RadioMode_CW || r.txMode == RadioMode_CWR {
			logrus.Infof("CW TX finished -- Override Freq")
			r.rig.SetMode(hamlib.VFOCurrent, hamlib.ModeLSB, -1)
			r.rig.SetFreq(hamlib.VFOCurrent, float64(r.txFreq+1500))
		}
		r.rig.SetPtt(hamlib.VFOCurrent, 1)
		r.audioTx.Start()
		r.lock.Unlock()
	} else {
		logrus.Errorf("no audio device for TX")
	}
}
