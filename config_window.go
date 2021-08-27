package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"matsu.dev/toe-log/state"
)

var (
	configDialog *gtk.Dialog
)

func InitConfigDialog(builder *gtk.Builder) {
	configDialog = mustGetObj(builder, "config-dialog").(*gtk.Dialog)
}

func ShowConfigDialog() {
	glib.IdleAdd(func() {
		InitConfigDialog(state.GetState().Gui)
		configDialog.Show()
	})
}
