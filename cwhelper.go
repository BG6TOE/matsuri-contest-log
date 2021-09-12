package main

import (
	"strings"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	"go.bug.st/serial"
	"matsu.dev/matsuri-contest-log/hamlib"
	"matsu.dev/matsuri-contest-log/morsecode"
	"matsu.dev/matsuri-contest-log/state"
)

const (
	CWHELPER_CW_USE_NONE = iota
	CWHELPER_CW_USE_RTS
	CWHELPER_CW_USE_DTR
)

type CWHelper struct {
	lock sync.Mutex
	port serial.Port

	isSending     bool
	shouldSending bool
	wpm           int

	CWKeyUse int
	PTTUse   int
}

var cwSender CWHelper

func (w *CWHelper) Init(conf *state.RigConfig) {
	w.shouldSending = false
	w.isSending = false
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.port != nil {
		w.port.Close()
	}

	portMode := serial.Mode{}
	portMode.BaudRate = conf.Baudrate
	portMode.DataBits = conf.Databits

	switch conf.Parity {
	case int(hamlib.ParityEven):
		portMode.Parity = serial.EvenParity
	case int(hamlib.ParityOdd):
		portMode.Parity = serial.OddParity
	case int(hamlib.ParityNone):
		portMode.Parity = serial.NoParity
	}

	switch conf.Stopbits {
	case 1:
		portMode.StopBits = serial.OneStopBit
	case 2:
		portMode.StopBits = serial.TwoStopBits
	default:
		portMode.StopBits = serial.OneStopBit
	}

	w.CWKeyUse = conf.CWPin
	w.PTTUse = conf.PTTPin

	var err error
	w.port, err = serial.Open(conf.Portname, &portMode)
	if err != nil {
		logrus.Errorf("Failed to open serial port for PTT: %v", err)
	}
	w.port.SetRTS(false)
	w.port.SetDTR(false)
	w.wpm = 20
	logrus.Infof("Opened serial port for PTT")
}

func (w *CWHelper) EnableSendMorse() {
	w.shouldSending = true
}

func (w *CWHelper) DisableSendMorse() {
	w.shouldSending = false
}

func (w *CWHelper) SendMorse(sequence string) {
	w.lock.Lock()
	defer w.lock.Unlock()

	if w.port == nil {
		return
	}

	w.isSending = true
	defer func() { w.isSending = false }()

	var ptt func(bool) error
	var cw func(bool) error

	if w.CWKeyUse == CWHELPER_CW_USE_DTR {
		cw = w.port.SetDTR
	} else if w.CWKeyUse == CWHELPER_CW_USE_RTS {
		cw = w.port.SetRTS
	} else {
		return
	}

	if w.PTTUse == CWHELPER_CW_USE_DTR {
		ptt = w.port.SetDTR
	} else if w.PTTUse == CWHELPER_CW_USE_RTS {
		ptt = w.port.SetRTS
	} else {
		ptt = nil
	}

	defer func() {
		if ptt != nil {
			ptt(false)
		}
	}()
	defer cw(false)

	sendSequence := make([]int, 0)

	elementNano := 60 * 1000000000 / (w.wpm * 50)

	s := strings.ToLower(sequence)
	for _, r := range s {
		code := morsecode.GetCode(r)
		for _, v := range code {
			sendSequence = append(sendSequence, v)
			if v > 0 {
				sendSequence = append(sendSequence, -1)
			}
		}
		sendSequence = append(sendSequence, -3)
	}

	if ptt != nil {
		ptt(true)
	}

	time.Sleep(50 * time.Millisecond)

	curr := 0
	total := len(sendSequence)
	nextWake := int64(0)
	for w.isSending && w.shouldSending && curr < total {
		if sendSequence[curr] > 0 {
			nextWake = time.Now().UnixNano() + int64(elementNano)*int64(sendSequence[curr])
		} else {
			nextWake = time.Now().UnixNano() - int64(elementNano)*int64(sendSequence[curr])
		}
		cw(sendSequence[curr] > 0)
		curr++
		for time.Now().UnixNano() < nextWake {
		}
	}
}
