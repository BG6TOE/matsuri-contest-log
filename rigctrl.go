package main

import (
	"time"

	"github.com/dh1tw/goHamlib"
	"matsu.dev/toe-log/state"
)

func InitRigctrl() {
	rig := state.GetState().HamlibRig
	rig.Init(2)
	rig.SetPort(goHamlib.Port{
		RigPortType: goHamlib.RigPortNetwork,
		Portname:    "127.0.0.1:4532",
	})
	rig.Open()

	go func() {
		for {
			freq, _ := rig.GetFreq(goHamlib.VFOCurrent)
			newVFO := uint64(freq)
			if newVFO != state.GetState().Rig.VFO {
				state.GetState().Rig.VFO = uint64(freq)
				state.Kick()
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func SetRigActiveFreq(freq float64) {
	rig := state.GetState().HamlibRig
	rig.SetFreq(goHamlib.VFOCurrent, float64(freq))
}
