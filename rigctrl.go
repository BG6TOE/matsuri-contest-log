package main

import (
	"time"

	"github.com/dh1tw/goHamlib"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	resources "matsu.dev/toe-log/resources"
	"matsu.dev/toe-log/state"
)

var (
	isRigReady = false
)

func InitRigctrl() {
	go ResetRig()
	go RefreshFreq()
}

func RefreshFreq() {
	rig := state.GetState().HamlibRig
	for {
		if isRigReady {
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
		return
	}

	setRadioStatusLight(resources.StatusLightIdle)

	isRigReady = true
}
