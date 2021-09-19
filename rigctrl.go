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
	"matsu.dev/matsuri-contest-log/webui/ws"
)

var (
	isRigReady = false
	rigMutex   = sync.Mutex{}
)

const (
	maxErrorsBeforeShuttingDownRig = 8
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
	totalError := 0
	rig := state.GetState().HamlibRig
	for {
		if isRigReady {
			func() {
				rigMutex.Lock()
				defer rigMutex.Unlock()

				needKick := false

				freq, err := rig.GetFreq(goHamlib.VFOCurrent)
				if err != nil {
					totalError++
				} else {
					totalError = 0
				}
				newVFO := uint64(freq)
				mode, _, err := rig.GetMode(goHamlib.VFOCurrent)
				if err != nil {
					totalError++
				} else {
					totalError = 0
				}

				if mode != goHamlib.ModeNONE {
					if state.GetState().HamlibRigMode != mode {
						state.GetState().Rig.Mode = goHamlib.ModeName[mode]
						state.GetState().HamlibRigMode = mode
						needKick = true
					}
				}

				if newVFO != state.GetState().Rig.VFO {
					state.GetState().Rig.VFO = uint64(freq)
					needKick = true
				}

				if needKick {
					state.Kick()
				}

				ptt, _ := rig.GetPtt(goHamlib.VFOCurrent)
				if ptt != goHamlib.RIG_PTT_OFF {
					setRadioStatusLight(resources.StatusLightBusy)
				} else {
					setRadioStatusLight(resources.StatusLightIdle)
				}

				if totalError > maxErrorsBeforeShuttingDownRig {
					go ResetRig()
				}
			}()
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
		label := mustGetObj(state.GetState().Gui, "statuslight-radio").(*gtk.Button)
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

func ShutdownRig() {
	rig := state.GetState().HamlibRig

	rigMutex.Lock()
	defer rigMutex.Unlock()

	if isRigReady {
		rig.Close()
		rig.Cleanup()
	}

	setRadioStatusLight(resources.StatusLightDisabled)
	isRigReady = false

	ws.Broadcast(&ws.BroadcastMessage{
		Class: "RigShutdown",
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

	isRigReady = false

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

	cwSender.Init(rigConfig)

	isRigReady = true

	ws.Broadcast(&ws.BroadcastMessage{
		Class: "RigReady",
	})
}
