package main

import (
	"context"
	"fmt"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"
	logdb "matsu.dev/matsuri-contest-log/logdb"
	"matsu.dev/matsuri-contest-log/resources"
	"matsu.dev/matsuri-contest-log/state"
	"matsu.dev/matsuri-contest-log/webui/ws"
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
		DXCallsign: callsign,
		OpCallsign: state.GetState().Operator.Callsign,
		FreqHz:     state.GetState().Rig.VFO,
		Mode:       state.GetState().Rig.Mode,
		RSTSent:    rstSent,
		RSTRcvd:    rstRcvd,
		ExchSent:   exchSent,
		ExchRcvd:   exchRcvd,
	}

	err := logdb.NewQSO(&state.GetState().Contest, &QSO)
	if err != nil {
		emitInfomation(builder, fmt.Sprintf("Failed to add QSO: %v", err), resources.InfoClassError)
		return
	}
	callsignEntry.SetText("")
	inputExchRcvdEntry.SetText("")
	emitInfomation(builder, fmt.Sprintf("Added QSO with %s", callsign), resources.InfoClassNotice)
}

func updateUIFromState() {
	states := state.GetState()
	vfo := states.Rig.VFO
	contest := states.Contest.Name
	mode := states.Rig.Mode
	glib.IdleAdd(func() {
		window := mustGetObj(states.Gui, "main_window").(*gtk.ApplicationWindow)
		window.SetTitle(fmt.Sprintf("%s: %s %d.%03d kHz", contest, mode, vfo/1000, vfo%1000))

		mustGetObj(states.Gui, "station-callsign").(*gtk.Label).SetText("Call: " + state.GetState().Contest.StationCallsign)
		mustGetObj(states.Gui, "operator-callsign").(*gtk.Label).SetText("Op: " + state.GetState().Operator.Callsign)
	})
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
	inputCallsign.ConnectAfter("changed", func(entry *gtk.Entry) {
		capitalizeInputHandler(entry)
		ws.Broadcast(&ws.BroadcastMessage{
			Class:   "InputCallsignChanged",
			Message: entry.GetChars(0, -1),
		})
	})

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
		currentText := entry.GetChars(0, -1)
		freq := float64(0)
		match, _ := regexp.MatchString(`^\d+(.(\d+)?)?$`, currentText)
		if match {
			fmt.Sscanf(currentText, "%f", &freq)
			if freq < 100 || freq >= 440*1000 {
				emitInfomation(builder, fmt.Sprintf("Frequency out of range: %f", freq), resources.InfoClassError)
				return
			}
			SetRigActiveFreq(freq * 1000)
			inputCallsign.SetText("")
			return
		}
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

func SetupMenuItems(builder *gtk.Builder, host gtk.IWindow) {
	{
		item := mustGetObj(builder, "menu-operation-config").(*gtk.MenuItem)
		item.Connect("activate", func() {
			configWindow.Show()
		})
	}
	{
		item := mustGetObj(builder, "menu-operation-export-adif").(*gtk.MenuItem)
		item.Connect("activate", func() {
			if userContinue, fileName := saveFileChooser(builder, host, "Export ADIF file to...", []string{"*.adif", "*.adi", "*.txt"}); userContinue {
				go func(builder *gtk.Builder, exportFileName string) {
					defaultContest := logdb.GetDefaultContext()
					if err := logdb.ExportADIF(&defaultContest, exportFileName); err != nil {
						emitInfomation(builder, fmt.Sprintf("Failed to export ADIF file: %v", err), resources.InfoClassError)
						return
					}
					emitInfomation(builder, fmt.Sprintf("Exported ADIF file: %v", exportFileName), resources.InfoClassNotice)
				}(builder, fileName)
			}
		})
	}
}

func saveFileChooser(builder *gtk.Builder, parent gtk.IWindow, title string, fileFilter []string) (bool, string) {
	dialog, _ := gtk.FileChooserDialogNewWith1Button(title, parent, gtk.FILE_CHOOSER_ACTION_SAVE, "Save", gtk.RESPONSE_ACCEPT)
	defer dialog.Destroy()
	dialog.SetDoOverwriteConfirmation(true)

	if len(fileFilter) > 0 {
		gtkFileFilter, err := gtk.FileFilterNew()
		if gtkFileFilter != nil {
			for _, item := range fileFilter {
				gtkFileFilter.AddPattern(item)
			}
			dialog.SetFilter(gtkFileFilter)
		} else {
			logrus.Warnf("Failed to init file filter: %s", err)
		}
	}
	resp := dialog.Run()
	if resp != gtk.RESPONSE_ACCEPT {
		logrus.Info("User cancelled")
		return false, ""
	}
	return true, dialog.GetFilename()
}

func SendCWMacro(key uint) {
	templateCtx := &MacroRenderContext{
		MyCall:  state.GetState().Contest.StationCallsign,
		DxCall:  mustGetObj(state.GetState().Gui, "input-callsign").(*gtk.Entry).GetChars(0, -1),
		Exch:    mustGetObj(state.GetState().Gui, "input-exch-sent").(*gtk.Entry).GetChars(0, -1),
		ExchRST: mustGetObj(state.GetState().Gui, "input-rst-sent").(*gtk.Entry).GetChars(0, -1),
	}

	cwToSent := ""
	render := macroRenderer.RenderMacroKey
	switch key {
	case gdk.KEY_F1:
		cwToSent = render("F1", templateCtx)
	case gdk.KEY_F2:
		cwToSent = render("F2", templateCtx)
	case gdk.KEY_F3:
		cwToSent = render("F3", templateCtx)
	case gdk.KEY_F4:
		cwToSent = render("F4", templateCtx)
	case gdk.KEY_F5:
		cwToSent = render("F5", templateCtx)
	case gdk.KEY_F6:
		cwToSent = render("F6", templateCtx)
	case gdk.KEY_F7:
		cwToSent = render("F7", templateCtx)
	case gdk.KEY_F8:
		cwToSent = render("F8", templateCtx)
	case gdk.KEY_F9:
		cwToSent = render("F9", templateCtx)
	case gdk.KEY_F10:
		cwToSent = render("F10", templateCtx)
	case gdk.KEY_F11:
		cwToSent = render("F11", templateCtx)
	case gdk.KEY_F12:
		cwToSent = render("F12", templateCtx)
	}
	if cwToSent == "" {
		return
	}
	emitInfomation(state.GetState().Gui, fmt.Sprintf("Sending: %v", cwToSent), resources.InfoClassNotice)
	go func() {
		if runtime.GOOS == "windows" {
			ShutdownRig()
			time.Sleep(50 * time.Millisecond)
			if len(state.GetState().RigConfig) > 0 {
				cwSender.Init(&state.GetState().RigConfig[0])
			}
			setRadioStatusLight(resources.StatusLightBusy)
		}
		cwSender.EnableSendMorse()
		cwSender.SendMorse(cwToSent)
		if runtime.GOOS == "windows" {
			cwSender.Deinit()
			time.Sleep(50 * time.Millisecond)
			ResetRig()
		}
	}()
}

func InitMacroKey(builder *gtk.Builder) {
	getMacroKeyConfig := func(key string) (state.MacroKeyMap, bool) {
		for _, v := range state.GetState().MacroKeyMap {
			if v.Key == key {
				return v, true
			}
		}
		return state.MacroKeyMap{}, false
	}

	for i := 1; i <= 12; i++ {
		key, ok := getMacroKeyConfig(fmt.Sprintf("F%d", i))
		btn := mustGetObj(builder, fmt.Sprintf("f%d-button", i)).(*gtk.Button)

		btn.SetLabel(fmt.Sprintf("F%d %s", i, key.Title))
		if !ok || key.Value == "" {
			btn.SetSensitive(false)
			continue
		}
		btn.SetSensitive(true)
	}
}

func InitMainWindow(builder *gtk.Builder, application *gtk.Application) {
	win := mustGetObj(builder, "main_window").(*gtk.ApplicationWindow)
	win.SetTitlebar(mustGetObj(builder, "application-header-bar").(*gtk.HeaderBar))

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

	win.SetTitle("Matsuri Contest Log")

	appContext, cancelFunc := context.WithCancel(context.Background())
	win.Connect("destroy", func() {
		cancelFunc()
		gtk.MainQuit()
	})

	InitMacroKey(builder)

	func() {
		btns := []uint{
			0,
			gdk.KEY_F1,
			gdk.KEY_F2,
			gdk.KEY_F3,
			gdk.KEY_F4,
			gdk.KEY_F5,
			gdk.KEY_F6,
			gdk.KEY_F7,
			gdk.KEY_F8,
			gdk.KEY_F9,
			gdk.KEY_F10,
			gdk.KEY_F11,
			gdk.KEY_F12,
		}
		for i := 1; i <= 12; i++ {
			func(btnid int) {
				btn := mustGetObj(builder, fmt.Sprintf("f%d-button", i)).(*gtk.Button)
				btn.Connect("clicked", func(btn *gtk.Button) {
					SendCWMacro(btns[btnid])
				})
			}(i)
		}
	}()

	win.Connect("key_press_event", func(win *gtk.ApplicationWindow, event *gdk.Event) {
		key := gdk.EventKeyNewFromEvent(event)
		if key.KeyVal() == gdk.KEY_F1 && ((key.State() & gdk.CONTROL_MASK) != 0) {
			configWindow.Show()
			return
		}
		if key.KeyVal() == gdk.KEY_Escape {
			cwSender.DisableSendMorse()
			return
		}

		switch key.KeyVal() {
		case gdk.KEY_F1:
			fallthrough
		case gdk.KEY_F2:
			fallthrough
		case gdk.KEY_F3:
			fallthrough
		case gdk.KEY_F4:
			fallthrough
		case gdk.KEY_F5:
			fallthrough
		case gdk.KEY_F6:
			fallthrough
		case gdk.KEY_F7:
			fallthrough
		case gdk.KEY_F8:
			fallthrough
		case gdk.KEY_F9:
			fallthrough
		case gdk.KEY_F10:
			fallthrough
		case gdk.KEY_F11:
			fallthrough
		case gdk.KEY_F12:
			SendCWMacro(key.KeyVal())
		}
	})

	win.ShowAll()

	utcClockLabel := mustGetObj(builder, "utc-clock").(*gtk.Label)

	current := time.Now().UTC()
	utcClockLabel.SetLabel(fmt.Sprintf("%02d:%02d:%02dZ", current.Hour(), current.Minute(), current.Second()))
	mustGetObj(builder, "station-callsign").(*gtk.Label).SetText("Call: " + state.GetState().Contest.StationCallsign)
	mustGetObj(builder, "operator-callsign").(*gtk.Label).SetText("Op: " + state.GetState().Operator.Callsign)

	mustGetObj(state.GetState().Gui, "statuslight-radio").(*gtk.Button).Connect("clicked", func(btn *gtk.Button) {
		go func() {
			glib.IdleAdd(
				func() {
					btn.SetSensitive(false)
				})
			ShutdownRig()
			time.Sleep(1 * time.Second)
			ResetRig()
			glib.IdleAdd(
				func() {
					btn.SetSensitive(true)
				})
		}()
	})

	state.RegisterStateChangeCallback(updateUIFromState)
	state.SetupStateTick()
	go timer(appContext, utcClockLabel)

	setupInputEvents(builder)
	SetupMenuItems(builder, win)
}
