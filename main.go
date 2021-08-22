package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"

	resources "matsu.dev/toe-log/resources"
)

func timer(ctx context.Context, timeLabel *gtk.Label) {
	loop := true
	for loop {
		select {
		case <-ctx.Done():
			loop = false
		case <-time.After(1 * time.Second):
			glib.IdleAdd(func() {
				current := time.Now().UTC()
				timeLabel.SetLabel(fmt.Sprintf("%02d:%02d:%02dZ", current.Hour(), current.Minute(), current.Second()))
			})
		}
	}
	logrus.Infof("App exit, stop timer goroutine")
}

func mustGetObj(builder *gtk.Builder, id string) glib.IObject {
	obj, err := builder.GetObject(id)
	if err != nil {
		logrus.Panicf("cannot get %s: %v", id, err)
	}
	return obj
}

func setupInputEvents(builder *gtk.Builder) {
	capitalizeInputHandler := func(editable *gtk.Entry) {
		buf := editable.GetChars(0, -1)
		capBuf2 := strings.Builder{}
		filtered := false
		for _, r := range buf {
			if (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || (r == '.') || (r == '/') || (r == '_') {
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
	numericInputHandler := func(editable *gtk.Entry) {
		buf := editable.GetChars(0, -1)
		capBuf2 := strings.Builder{}
		filtered := false
		for _, r := range buf {
			if r >= '0' && r <= '9' {
				capBuf2.WriteRune(r)
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

	inputCallsign := mustGetObj(builder, "input-callsign").(*gtk.Entry)
	inputCallsign.ConnectAfter("changed", capitalizeInputHandler)

	inputExchSent := mustGetObj(builder, "input-exch-sent").(*gtk.Entry)
	inputExchSent.ConnectAfter("changed", capitalizeInputHandler)

	inputExchRcvd := mustGetObj(builder, "input-exch-rcvd").(*gtk.Entry)
	inputExchRcvd.ConnectAfter("changed", capitalizeInputHandler)

	inputRstSent := mustGetObj(builder, "input-rst-sent").(*gtk.Entry)
	inputRstSent.ConnectAfter("changed", numericInputHandler)

	inputRstRcvd := mustGetObj(builder, "input-rst-rcvd").(*gtk.Entry)
	inputRstRcvd.ConnectAfter("changed", numericInputHandler)
}

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	builder, err := gtk.BuilderNewFromString(resources.Glade)
	if err != nil {
		logrus.Fatalf("Unable to init builder: %v", err)
	}

	obj, err := builder.GetObject("main_window")
	if err != nil {
		logrus.Fatalf("Failed to get main_window")
	}
	win := obj.(*gtk.Window)

	appContext := context.Background()
	appContextWithCancel, cancelFunc := context.WithCancel(appContext)
	_ = appContextWithCancel
	if err != nil {
		logrus.Fatal("Unable to create window:", err)
	}

	{
		css, err := gtk.CssProviderNew()
		if err != nil {
			logrus.Fatalf("Unable to init css provider: %v", err)
		}

		err = css.LoadFromData(resources.CSS)
		if err != nil {
			logrus.Fatalf("Unable to init css: %v", err)
		}

		screen := win.GetScreen()
		gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_USER)
	}

	win.SetTitle("TOE Contest Log")
	win.Connect("destroy", func() {
		cancelFunc()
		gtk.MainQuit()
	})

	win.ShowAll()

	utcClockLabel := mustGetObj(builder, "utc-clock").(*gtk.Label)

	current := time.Now().UTC()
	utcClockLabel.SetLabel(fmt.Sprintf("%02d:%02d:%02dZ", current.Hour(), current.Minute(), current.Second()))

	go timer(appContextWithCancel, utcClockLabel)

	setupInputEvents(builder)

	gtk.Main()
}
