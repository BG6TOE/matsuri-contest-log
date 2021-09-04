package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
	goHamlib "matsu.dev/matsuri-contest-log/hamlib"
	resources "matsu.dev/matsuri-contest-log/resources"
	"matsu.dev/matsuri-contest-log/state"
)

var (
	isRigReady = false
	rigMutex   = sync.Mutex{}
)

func init() {
	goHamlib.SetDebugLevel(goHamlib.DebugWarn)
	goHamlib.SetDebugCallback(func(lvl goHamlib.DebugLevel, text string) {
		switch lvl {
		case goHamlib.DebugNone:
		case goHamlib.DebugBug:
		case goHamlib.DebugErr:
			logrus.Errorf("hamlib: %v", text)
		case goHamlib.DebugWarn:
			logrus.Warnf("hamlib: %v", text)
		case goHamlib.DebugVerbose:
			logrus.Infof("hamlib: %v", text)
		case goHamlib.DebugTrace:
			logrus.Debugf("hamlib: %v", text)
		}
	})
}

func InitRigctrl() {
	go ResetRig()
	go RefreshFreq()
}

func RefreshFreq() {
	rig := state.GetState().HamlibRig
	for {
		if isRigReady {
			func() {
				rigMutex.Lock()
				defer rigMutex.Unlock()

				freq, _ := rig.GetFreq(goHamlib.VFOCurrent)
				newVFO := uint64(freq)
				if newVFO != state.GetState().Rig.VFO {
					state.GetState().Rig.VFO = uint64(freq)
					state.Kick()
				}

				ptt, _ := rig.GetPtt(goHamlib.VFOCurrent)
				if ptt != goHamlib.RIG_PTT_OFF {
					setRadioStatusLight(resources.StatusLightBusy)
				} else {
					setRadioStatusLight(resources.StatusLightIdle)
				}
			}()
		} else {
			state.GetState().Rig.VFO = 0
		}
		time.Sleep(1 * time.Second)
	}
}

func SetRigActiveFreq(freq float64) {
	rig := state.GetState().HamlibRig
	rig.SetFreq(goHamlib.VFOCurrent, float64(freq))
}

func setRadioStatusLight(signal string) {
	glib.IdleAdd(func() {
		label := mustGetObj(state.GetState().Gui, "statuslight-radio").(*gtk.Label)
		styleCtx, err := label.GetStyleContext()
		if err != nil {
			return
		}
		for _, v := range resources.StatusLightClasses {
			styleCtx.RemoveClass(v)
		}
		styleCtx.AddClass(signal)
	})
}

func ResetRig() {
	rig := state.GetState().HamlibRig

	rigMutex.Lock()
	defer rigMutex.Unlock()

	if isRigReady {
		rig.Close()
		rig.Cleanup()
	}

	if len(state.GetState().RigConfig) == 0 {
		return
	}

	rigConfig := &state.GetState().RigConfig[0]

	setRadioStatusLight(resources.StatusLightInit)

	ifFindRigModel := false
	for _, v := range goHamlib.ListModels() {
		if v.Manufacturer == rigConfig.Manufacturer && v.Model == rigConfig.Model {
			rig.Init(v.ModelID)
			ifFindRigModel = true
			break
		}
	}

	if !ifFindRigModel {
		setRadioStatusLight(resources.StatusLightDisabled)
		return
	}

	var ok bool
	portConfig := goHamlib.Port{}
	if portConfig.RigPortType, ok = goHamlib.RigPortValue[rigConfig.RigPortType]; !ok {
		setRadioStatusLight(resources.StatusLightDisabled)
		return
	}
	portConfig.Baudrate = rigConfig.Baudrate
	portConfig.Databits = rigConfig.Databits
	portConfig.Stopbits = rigConfig.Stopbits
	portConfig.Handshake = goHamlib.Handshake(rigConfig.Flowcontrol)
	portConfig.Parity = goHamlib.Parity(rigConfig.Parity)
	portConfig.Portname = rigConfig.Portname

	rig.SetPort(portConfig)

	if err := rig.Open(); err != nil {
		setRadioStatusLight(resources.StatusLightError)
		emitInfomation(state.GetState().Gui, fmt.Sprintf("Failed to open rig: %v", err), resources.InfoClassError)
		return
	}

	emitInfomation(state.GetState().Gui, "Connected to radio", resources.InfoClassNotice)
	setRadioStatusLight(resources.StatusLightIdle)

	isRigReady = true
}
