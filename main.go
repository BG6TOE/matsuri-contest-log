package main

import (
	"context"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"

	logdb "matsu.dev/toe-log/logdb"
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

func mustGetEntryBuf(entry *gtk.Entry) *gtk.EntryBuffer {
	buf, err := entry.GetBuffer()
	if err != nil {
		logrus.Panicf("cannot get buffer for entry!")
	}
	return buf
}

func emitInfomation(builder *gtk.Builder, info, infoClass string) {
	glib.IdleAdd(func() {
		infoLabel := mustGetObj(builder, "infomation").(*gtk.Label)
		styleCtx, err := infoLabel.GetStyleContext()
		if err != nil {
			return
		}
		styleCtx.RemoveClass(resources.InfoClassNotice)
		styleCtx.RemoveClass(resources.InfoClassWarning)
		styleCtx.RemoveClass(resources.InfoClassError)
		styleCtx.AddClass(infoClass)

		infoLabel.SetText(info)
	})
}

func commitQSO(builder *gtk.Builder) {
	callsignEntry := mustGetObj(builder, "input-callsign").(*gtk.Entry)
	rstSentEntry := mustGetObj(builder, "input-rst-sent").(*gtk.Entry)
	rstRcvdEntry := mustGetObj(builder, "input-rst-rcvd").(*gtk.Entry)
	inputExchSentEntry := mustGetObj(builder, "input-exch-sent").(*gtk.Entry)
	inputExchRcvdEntry := mustGetObj(builder, "input-exch-rcvd").(*gtk.Entry)

	callsign := callsignEntry.GetChars(0, -1)
	rstSent := rstSentEntry.GetChars(0, -1)
	rstRcvd := rstRcvdEntry.GetChars(0, -1)
	exchSent := inputExchSentEntry.GetChars(0, -1)
	exchRcvd := inputExchRcvdEntry.GetChars(0, -1)

	if callsign == "" {
		emitInfomation(builder, fmt.Sprintf("%s is not a valid callsign!", callsign), resources.InfoClassError)
		return
	}

	QSO := logdb.QSO{
		DXCallsign: callsign, //      string    `json:"dx_callsign"`
		FreqHz:     14014000, //     uint64    `json:"freq_hz"`
		Mode:       "CW",     //            string    `json:"mode"`
		RSTSent:    rstSent,  //         string    `json:"rst_sent"`
		RSTRcvd:    rstRcvd,  //         string    `json:"rst_rcvd"`
		ExchSent:   exchSent, //        string    `json:"exch_sent"`
		ExchRcvd:   exchRcvd, //        string    `json:"exch_rcvd"`
	}

	err := logdb.NewQSO(logdb.GetDefaultContext(), &QSO)
	if err != nil {
		emitInfomation(builder, fmt.Sprintf("Failed to add QSO: %v", err), resources.InfoClassError)
		return
	}
	emitInfomation(builder, fmt.Sprintf("Success added QSO with %s", callsign), resources.InfoClassNotice)
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

	setDefaultRST := func() {
		if inputRstRcvd.GetChars(0, -1) == "" {
			editableBuf, err := inputRstRcvd.GetBuffer()
			if err != nil {
				return
			}
			editableBuf.SetText("599")
		}
		if inputRstSent.GetChars(0, -1) == "" {
			editableBuf, err := inputRstSent.GetBuffer()
			if err != nil {
				return
			}
			editableBuf.SetText("599")
		}
	}

	inputCallsign.Connect("key-press-event", func(entry *gtk.Entry, event *gdk.Event) {
		key := gdk.EventKeyNewFromEvent(event)
		keyVal := key.KeyVal()
		if keyVal == gdk.KEY_space {
			inputExchRcvd.ToWidget().GrabFocus()
			textLen := int(inputExchRcvd.GetTextLength())
			inputExchRcvd.SelectRegion(textLen, textLen+1)
		} else {
			return
		}
		setDefaultRST()
	})

	inputCallsign.Connect("activate", func(entry *gtk.Entry) {
		setDefaultRST()
		inputExchRcvd.ToWidget().GrabFocus()
		textLen := int(inputExchRcvd.GetTextLength())
		inputExchRcvd.SelectRegion(textLen, textLen+1)
	})

	inputExchRcvd.Connect("key-press-event", func(entry *gtk.Entry, event *gdk.Event) {
		key := gdk.EventKeyNewFromEvent(event)
		if key.KeyVal() == gdk.KEY_space {
			inputCallsign.ToWidget().GrabFocus()
			textLen := int(inputCallsign.GetTextLength())
			inputCallsign.SelectRegion(textLen, textLen+1)
		}
		setDefaultRST()
	})

	inputExchRcvd.Connect("activate", func(entry *gtk.Entry) {
		commitQSO(builder)
	})
}

func initDatabase() {
	homdDir := glib.GetHomeDir()
	databaseDir := path.Join(homdDir, ".config", "dev.matsu.contestlog")
	databaseName := path.Join(databaseDir, "log.sqlite3")
	if err := os.MkdirAll(databaseDir, os.ModeDir); err != nil {
		logrus.Fatalf("Failed to create directory: %v", err)
	}
	if err := logdb.Init(databaseName); err != nil {
		logrus.Fatalf("Failed to setup main database: %v", err)
	}
}

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	initDatabase()

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
