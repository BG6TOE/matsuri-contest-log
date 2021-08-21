package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
)

func timer(ctx context.Context, timeLabel *gtk.Label) {
	loop := true
	for loop {
		select {
		case <-ctx.Done():
			loop = false
		case <-time.After(1 * time.Second):
			current := time.Now().UTC()
			timeLabel.SetLabel(fmt.Sprintf("%02d:%02d:%02dZ", current.Hour(), current.Minute(), current.Second()))
		}
	}
	logrus.Infof("App exit, stop timer goroutine")
}

func tryAddClassToWidget(w gtk.Widget, class string) {
	styleContext, err := w.GetStyleContext()
	if err != nil {
		logrus.Error("Failed to get style context")
		return
	}
	styleContext.AddClass(class)
}

func mustNewLabel(initText string) *gtk.Label {
	lab, err := gtk.LabelNew(initText)
	if err != nil {
		logrus.Fatalf("Failed to create new GtkLabel: %v", err)
	}
	return lab
}

func mustNewEntry() *gtk.Entry {
	entry, err := gtk.EntryNew()
	if err != nil {
		logrus.Fatalf("Failed to create new GtkEntry: %v", err)
	}
	return entry
}

func mustSetupInputGrid(ctx context.Context) (inputGrid *gtk.Grid) {
	inputGrid, err := gtk.GridNew()
	if err != nil {
		logrus.Fatalf("Failed to setup input grid: %v", err)
	}

	hintLabelCallsign := mustNewLabel("Callsign")
	hintLabelSentRST := mustNewLabel("RST Sent")
	hintLabelRcvdRST := mustNewLabel("RST Rcvd")
	hintLabelSentExch := mustNewLabel("Exch Sent")
	hintLabelRcvdExch := mustNewLabel("Exch Rcvd")

	inputCallsign := mustNewEntry()
	inputSentRST := mustNewEntry()
	inputRcvdRST := mustNewEntry()
	inputSentExch := mustNewEntry()
	inputRcvdExch := mustNewEntry()

	inputCallsign.SetName("callsign_input")
	inputSentRST.SetName("sent_rst")
	inputRcvdRST.SetName("rcvd_rst")
	inputSentExch.SetName("sent_exch")
	inputRcvdExch.SetName("rcvd_exch")

	mainInputFilter := func(editable *gtk.Entry) {
		buf := editable.GetChars(0, -1)
		capBuf2 := strings.Builder{}
		filtered := false
		strlen := 0
		for _, r := range buf {
			if strlen >= 10 {
				filtered = true
				continue
			}
			if (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') || (r == '.') || (r == '/') || (r == '_') {
				capBuf2.WriteRune(r)
				strlen++
			} else if r >= 'a' && r <= 'z' {
				capBuf2.WriteRune(r - 'a' + 'A')
				strlen++
				filtered = true
			} else {
				filtered = true
			}
		}
		if !filtered {
			return
		}
		newBuf, err := gtk.EntryBufferNew(capBuf2.String(), capBuf2.Len())
		if err != nil {
			return
		}
		editable.SetBuffer(newBuf)
	}

	rstInputFiler := func(editable *gtk.Entry) {
		buf := editable.GetChars(0, -1)
		capBuf2 := strings.Builder{}
		filtered := false
		strlen := 0
		for _, r := range buf {
			if strlen >= 3 {
				filtered = true
				continue
			}
			if r >= '0' && r <= '9' {
				capBuf2.WriteRune(r)
				strlen++
			} else {
				filtered = true
			}
		}
		if !filtered {
			return
		}
		newBuf, err := gtk.EntryBufferNew(capBuf2.String(), capBuf2.Len())
		if err != nil {
			return
		}
		editable.SetBuffer(newBuf)
	}

	inputCallsign.Connect("changed", mainInputFilter)
	inputRcvdRST.Connect("changed", rstInputFiler)
	inputSentRST.Connect("changed", rstInputFiler)
	inputRcvdExch.Connect("changed", mainInputFilter)
	inputSentExch.Connect("changed", mainInputFilter)

	inputGrid.Add(hintLabelCallsign)
	inputGrid.AttachNextTo(inputCallsign, hintLabelCallsign, gtk.POS_BOTTOM, 1, 1)
	inputGrid.AttachNextTo(hintLabelSentRST, hintLabelCallsign, gtk.POS_RIGHT, 1, 1)
	inputGrid.AttachNextTo(inputSentRST, hintLabelSentRST, gtk.POS_BOTTOM, 1, 1)
	inputGrid.AttachNextTo(hintLabelRcvdRST, hintLabelSentRST, gtk.POS_RIGHT, 1, 1)
	inputGrid.AttachNextTo(inputRcvdRST, hintLabelRcvdRST, gtk.POS_BOTTOM, 1, 1)
	inputGrid.AttachNextTo(hintLabelSentExch, hintLabelRcvdRST, gtk.POS_RIGHT, 1, 1)
	inputGrid.AttachNextTo(inputSentExch, hintLabelSentExch, gtk.POS_BOTTOM, 1, 1)
	inputGrid.AttachNextTo(hintLabelRcvdExch, hintLabelSentExch, gtk.POS_RIGHT, 1, 1)
	inputGrid.AttachNextTo(inputRcvdExch, hintLabelRcvdExch, gtk.POS_BOTTOM, 1, 1)

	hintLabelCallsign.SetHAlign(gtk.ALIGN_START)
	hintLabelSentRST.SetHAlign(gtk.ALIGN_START)
	hintLabelRcvdRST.SetHAlign(gtk.ALIGN_START)
	hintLabelSentExch.SetHAlign(gtk.ALIGN_START)
	hintLabelRcvdExch.SetHAlign(gtk.ALIGN_START)

	tryAddClassToWidget(hintLabelCallsign.Widget, "context-exch-hint-label")
	tryAddClassToWidget(hintLabelSentRST.Widget, "context-exch-hint-label")
	tryAddClassToWidget(hintLabelRcvdRST.Widget, "context-exch-hint-label")
	tryAddClassToWidget(hintLabelSentExch.Widget, "context-exch-hint-label")
	tryAddClassToWidget(hintLabelRcvdExch.Widget, "context-exch-hint-label")

	tryAddClassToWidget(inputCallsign.Widget, "context-exch")
	tryAddClassToWidget(inputSentRST.Widget, "context-exch")
	tryAddClassToWidget(inputRcvdRST.Widget, "context-exch")
	tryAddClassToWidget(inputSentExch.Widget, "context-exch")
	tryAddClassToWidget(inputRcvdExch.Widget, "context-exch")

	return
}

func mustSetupTitleGrid(ctx context.Context) (titleGrid *gtk.Grid) {
	titleGrid, err := gtk.GridNew()
	if err != nil {
		logrus.Fatal("Failed to setup grid for title: %v", err)
	}

	timeLabel := mustNewLabel("xx:xx:xxZ")
	timeLabel.SetHAlign(gtk.ALIGN_END)
	timeLabel.SetName("timerLabel")

	opLabel := mustNewLabel("OP: xx0xxx")
	opLabel.SetName("operator_name")

	callLabel := mustNewLabel("STA: xx0xxx")
	callLabel.SetHAlign(gtk.ALIGN_START)
	callLabel.SetName("operator_name")

	tryAddClassToWidget(timeLabel.Widget, "ui-header-label")
	tryAddClassToWidget(opLabel.Widget, "ui-header-label")
	tryAddClassToWidget(callLabel.Widget, "ui-header-label")

	titleGrid.Add(timeLabel)
	titleGrid.AttachNextTo(opLabel, timeLabel, gtk.POS_LEFT, 1, 1)
	titleGrid.AttachNextTo(callLabel, opLabel, gtk.POS_LEFT, 1, 1)

	go timer(ctx, timeLabel)

	return
}

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	appContext := context.Background()
	appContextWithCancel, cancelFunc := context.WithCancel(appContext)
	if err != nil {
		logrus.Fatal("Unable to create window:", err)
	}

	{
		css, err := gtk.CssProviderNew()
		if err != nil {
			logrus.Fatalf("Unable to init css provider: %v", err)
		}

		err = css.LoadFromData(`
		.context-exch { font-family: 'Fira Code', monospace; font-size: xx-large; }
		.context-exch-hint-label { font-family: 'Fira Code', monospace; font-size: large; }
		.ui-header-label { font-family: 'Fira Code', monospace; font-size: large; font-weight: 900; }
		`)
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

	mainGrid, err := gtk.GridNew()
	if err != nil {
		logrus.Fatal("Unable to create main grid:", err)
	}

	win.Add(mainGrid)
	mainGrid.SetHExpand(true)
	mainGrid.SetVExpand(true)

	titleGrid := mustSetupTitleGrid(appContextWithCancel)
	titleGrid.SetHExpand(true)
	titleGrid.SetColumnHomogeneous(true)
	mainGrid.Add(titleGrid)

	inputGrid := mustSetupInputGrid(appContextWithCancel)
	inputGrid.SetHExpand(true)
	inputGrid.SetColumnHomogeneous(true)
	inputGrid.SetColumnSpacing(5)

	mainGrid.AttachNextTo(inputGrid, titleGrid, gtk.POS_BOTTOM, 1, 1)
	mainGrid.SetRowSpacing(10)
	mainGrid.SetBorderWidth(10)

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	win.SetTitle("Station: xx0xxx  Op: xx0xxx")

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
