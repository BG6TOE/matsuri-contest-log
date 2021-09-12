package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
	"matsu.dev/matsuri-contest-log/state"
)

type macroConfig struct {
	macroTreeView  *gtk.TreeView
	macroListStore *gtk.ListStore
	macros         map[string]string

	defaultMacro []state.MacroKeyMap
}

const (
	CONFIG_MACRO_COLUMN_KEY = iota
	CONFIG_MACRO_COLUMN_TITLE
	CONFIG_MACRO_COLUMN_MACRO
)

func (mc *macroConfig) createColumn(title string, id int, editable bool) *gtk.TreeViewColumn {
	renderer, err := gtk.CellRendererTextNew()
	if err != nil {
		logrus.Fatal("Failed to create text cell renderer: ", err)
	}
	renderer.Set("editable", editable)
	col, err := gtk.TreeViewColumnNewWithAttribute(title, renderer, "text", id)
	if err != nil {
		logrus.Fatal("Failed to create tree view column: ", err)
	}
	return col
}

func (mc *macroConfig) insertMacroKeySetting(key, title, macro string) {
	if err := mc.macroListStore.Set(
		mc.macroListStore.Append(),
		[]int{CONFIG_MACRO_COLUMN_KEY, CONFIG_MACRO_COLUMN_TITLE, CONFIG_MACRO_COLUMN_MACRO},
		[]interface{}{key, title, macro}); err != nil {
		logrus.Fatalf("Failed to insert items: %v", err)
	}
}

func (mc *macroConfig) initMacroConfig(builder *gtk.Builder) {
	var err error

	mc.defaultMacro = []state.MacroKeyMap{
		{"F1", "CQ", "CQ {{.MyCall}}"},
		{"F2", "Exch", "{{.ExchRST}} {{.Exch}}"},
		{"F3", "TU", "TU"},
		{"F4", "My Call", "{{.MyCall}}"},
		{"F5", "DX Call", "{{.DxCall}}"},
		{"F6", "", ""},
		{"F7", "", ""},
		{"F8", "", ""},
		{"F9", "", ""},
		{"F10", "", ""},
		{"F11", "", ""},
		{"F12", "Clear", "_CLEAR"},
	}

	mc.macroTreeView = mustGetObj(builder, "config-macros").(*gtk.TreeView)
	mc.macroListStore, err = gtk.ListStoreNew(glib.TYPE_STRING, glib.TYPE_STRING, glib.TYPE_STRING)
	if err != nil {
		logrus.Fatalf("Failed to create a new store view: %v", err)
	}

	mc.macroTreeView.AppendColumn(mc.createColumn("Key", CONFIG_MACRO_COLUMN_KEY, false))
	mc.macroTreeView.AppendColumn(mc.createColumn("Name", CONFIG_MACRO_COLUMN_TITLE, true))
	mc.macroTreeView.AppendColumn(mc.createColumn("Macro", CONFIG_MACRO_COLUMN_MACRO, true))
	mc.macroTreeView.SetModel(mc.macroListStore.ToTreeModel())

	if state.GetState().MacroKeyMap == nil {
		state.GetState().MacroKeyMap = mc.defaultMacro
	}

	for _, v := range state.GetState().MacroKeyMap {
		mc.insertMacroKeySetting(v.Key, v.Title, v.Value)
	}
}
