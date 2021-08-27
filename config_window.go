package main

import (
	"fmt"
	"sort"

	"github.com/dh1tw/goHamlib"
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

	rigModelSelect := mustGetObj(builder, "config-rig-type").(*gtk.ComboBoxText)
	{
		rigModels := goHamlib.ListModels()
		rigModelSelect.RemoveAll()
		sort.Slice(rigModels, func(i, j int) bool {
			return rigModels[i].Manufacturer < rigModels[j].Manufacturer ||
				(rigModels[i].Manufacturer == rigModels[j].Manufacturer && rigModels[i].Model < rigModels[j].Model)
		})
		for _, v := range rigModels {
			rigModelSelect.Append(fmt.Sprintf("%d", v.ModelID), fmt.Sprintf("%s - %s (%d)", v.Manufacturer, v.Model, v.ModelID))
		}
	}

}

func ShowConfigDialog() {
	glib.IdleAdd(func() {
		configDialog.Show()
	})
}
