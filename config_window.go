package main

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"matsu.dev/matsuri-contest-log/state"
)

type ConfigWindow struct {
	contestConfig
	macroConfig

	dialog *gtk.Dialog
}

var configWindow ConfigWindow

func onSaveClicked(btn *gtk.Button) {
	configWindow.onSaveClickedRig(btn)
	configWindow.onSaveClickedContest(btn)

	state.SaveState()
}

func (c *ConfigWindow) Init(builder *gtk.Builder) {
	c.dialog = mustGetObj(builder, "config-dialog").(*gtk.Dialog)
	c.dialog.Connect("delete-event", func(dialog *gtk.Dialog, event *gdk.Event) bool {
		dialog.Hide()
		return true
	})
	c.initRigConfig(builder)
	c.initContestConfig(builder)
	c.initSystemConfig(builder)
	c.initMacroConfig(builder)
	mustGetObj(builder, "config-save").(*gtk.Button).Connect("clicked", onSaveClicked)
}

func (c *ConfigWindow) Show() {
	glib.IdleAdd(func() {
		c.dialog.Show()
	})
}

func (c *ConfigWindow) Hide() {
	glib.IdleAdd(func() {
		c.dialog.Hide()
	})
}
