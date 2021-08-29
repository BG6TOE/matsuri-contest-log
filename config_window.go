package main

import (
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

var (
	configDialog *gtk.Dialog
)

func InitConfigDialog(builder *gtk.Builder) {
	configDialog = mustGetObj(builder, "config-dialog").(*gtk.Dialog)
	configDialog.Connect("delete-event", func(dialog *gtk.Dialog, event *gdk.Event) bool {
		dialog.Hide()
		return true
	})
	InitConfigDialogRigConfig(builder)
}

func ShowConfigDialog() {
	glib.IdleAdd(func() {
		configDialog.Show()
	})
}

func HideConfigDialog() {
	glib.IdleAdd(func() {
		configDialog.Hide()
	})
}
