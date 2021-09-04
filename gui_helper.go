package main

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
)

func mustGetObj(builder *gtk.Builder, id string) glib.IObject {
	obj, err := builder.GetObject(id)
	if err != nil {
		logrus.Panicf("cannot get %s: %v", id, err)
	}
	return obj
}

func mustGetEntryBuf(entry *gtk.Entry) *gtk.EntryBuffer {
	buf, err := entry.GetBuffer()
	if err != nil {
		logrus.Panicf("cannot get buffer for entry!")
	}
	return buf
}
