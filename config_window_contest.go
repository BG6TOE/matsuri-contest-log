package main

import (
	"strings"

	"github.com/gotk3/gotk3/gtk"
	"matsu.dev/matsuri-contest-log/state"
)

type contestConfig struct {
	opCallsignEntry  *gtk.Entry
	staCallsignEntry *gtk.Entry
}

func (c *ConfigWindow) onSaveClickedContest(btn *gtk.Button) {
	state.GetState().Operator.Callsign = c.opCallsignEntry.GetChars(0, -1)
	state.GetState().Contest.StationCallsign = c.staCallsignEntry.GetChars(0, -1)

	state.Kick()
}

func (c *ConfigWindow) initContestConfig(builder *gtk.Builder) {
	operatorCallsign := state.GetState().Operator.Callsign
	stationCallsign := state.GetState().Contest.StationCallsign

	callsignInputHandler := func(editable *gtk.Entry) {
		buf := editable.GetChars(0, -1)
		capBuf2 := strings.Builder{}
		filtered := false
		for _, r := range buf {
			if (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || (r == '/') {
				capBuf2.WriteRune(r)
			} else if r >= 'a' && r <= 'z' {
				capBuf2.WriteRune(r - 'a' + 'A')
				filtered = true
			} else {
				filtered = true
			}
		}
		if !filtered {
			return
		}
		editableBuf, err := editable.GetBuffer()
		if err != nil {
			return
		}
		editableBuf.SetText(capBuf2.String())
	}

	c.opCallsignEntry = mustGetObj(builder, "config-operator-callsign").(*gtk.Entry)
	c.opCallsignEntry.SetText(operatorCallsign)
	c.opCallsignEntry.ConnectAfter("changed", callsignInputHandler)

	c.staCallsignEntry = mustGetObj(builder, "config-station-callsign").(*gtk.Entry)
	c.staCallsignEntry.SetText(stationCallsign)
	c.staCallsignEntry.ConnectAfter("changed", callsignInputHandler)
}
