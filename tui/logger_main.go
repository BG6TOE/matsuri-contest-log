package main

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	gproto "matsu.dev/matsuri-contest-log/proto"
)

var (
	activeContest *gproto.ActiveContest
	stopChan      chan interface{}

	mainPage *tview.Pages
)

func formatDateTime(dt time.Time) string {
	dt = dt.UTC()
	return fmt.Sprintf("%02d-%02d %02d:%02d", int(dt.Month()), int(dt.Day()), dt.Hour(), dt.Minute())
}

func formatFreq(freq int64) string {
	return fmt.Sprintf("%d.%d", freq/1000, (freq/100)%10)
}

func clockLoop(app *tview.Application, field *tview.InputField) {
	drawTime := func() {
		field.SetText(formatDateTime(time.Now()))
	}
	app.QueueUpdateDraw(drawTime)

for_loop:
	for {
		timeToNextMinute := 60 - time.Now().UTC().Second()%60
		select {
		case <-stopChan:
			break for_loop
		case <-time.After(time.Second * time.Duration(timeToNextMinute)):
			app.QueueUpdateDraw(drawTime)
		}
	}
}

type qsoPage struct {
	app *tview.Application

	currentQso    string
	fields        []*tview.InputField
	fieldsNameMap map[string]*tview.InputField

	qsoTable         *tview.Table
	callsignHintList *tview.TextView

	acceptFunctions    map[int]func(textToCheck string, lastChar rune) bool
	focusLossFunctions map[int]func(key tcell.Key)

	radioPanel *radioPanel
}

type radioPanel struct {
	freq   uint64
	mode   string
	signal string
	active bool

	configButton  *tview.Button
	freqField     *tview.InputField
	modeField     *tview.DropDown
	signalField   *tview.InputField
	radioDialogue *tview.Form

	radioFormModel *tview.DropDown
	radioFormUrl   *tview.InputField
}

var qsoPageData qsoPage

func createAcceptFunction(form *qsoPage, index int) func(textToCheck string, lastChar rune) bool {
	return func(textToCheck string, lastChar rune) bool {
		fn, ok := form.acceptFunctions[index]
		return !ok || fn(textToCheck, lastChar)
	}
}

func createFocusLossFunction(form *qsoPage, index int) func(key tcell.Key) {
	return func(key tcell.Key) {
		fn, ok := form.focusLossFunctions[index]
		if ok {
			fn(key)
		}
	}
}

func (f *qsoPage) addFieldWithName(name string, field *tview.InputField) {
	f.fields = append(f.fields, field)
	if len(name) > 0 {
		f.fieldsNameMap[name] = field
	}
}

func (f *qsoPage) getFieldValue(name string) string {
	field, ok := f.fieldsNameMap[name]
	if !ok {
		return ""
	}
	return field.GetText()
}

func (f *qsoPage) setFieldValue(name string, value string) {
	field, ok := f.fieldsNameMap[name]
	if !ok {
		return
	}
	field.SetText(value)
}

func (f *qsoPage) focusField(name string) {
	field, ok := f.fieldsNameMap[name]
	if !ok {
		return
	}
	f.app.SetFocus(field)
}

func (f *qsoPage) draftQso() {
	exchSent := make(map[string]string)
	exchRcvd := make(map[string]string)

	for _, customField := range activeContest.Contest.ExchangeSent {
		exchSent[customField] = f.getFieldValue(customField)
	}
	for _, customField := range activeContest.Contest.ExchangeRcvd {
		exchRcvd[customField] = f.getFieldValue(customField)
	}

	qso := &gproto.DraftQSOMessage{
		Uid:          f.currentQso,
		DxCallsign:   f.getFieldValue("dx_callsign"),
		Mode:         f.getFieldValue("mode"),
		ExchangeSent: exchSent,
		ExchangeRcvd: exchRcvd,
	}

	qso, err := realTimeGuiClient.DraftQSO(context.Background(), qso)
	if err != nil {
		return
	}

	f.app.QueueUpdateDraw(func() {
		f.currentQso = qso.Uid
		for k, v := range qso.ExchangeSent {
			f.setFieldValue(k, v)
		}
		for k, v := range qso.ExchangeRcvd {
			f.setFieldValue(k, v)
		}
		if qso.Expect != "" {
			f.focusField(qso.Expect)
		}
	})
}

func (f *qsoPage) submitQso() {
	exchSent := make(map[string]string)
	exchRcvd := make(map[string]string)

	for _, customField := range activeContest.Contest.ExchangeSent {
		exchSent[customField] = f.getFieldValue(customField)
	}
	for _, customField := range activeContest.Contest.ExchangeRcvd {
		exchRcvd[customField] = f.getFieldValue(customField)
	}

	qso := &gproto.QSO{
		DxCallsign:   f.getFieldValue("dx_callsign"),
		Time:         time.Now().Unix(),
		Freq:         int64(f.radioPanel.freq),
		Mode:         f.getFieldValue("mode"),
		ExchangeSent: exchSent,
		ExchangeRcvd: exchRcvd,
	}

	guiClient.LogQSO(context.Background(), qso)
	f.reloadQsoTable()

	f.app.QueueUpdateDraw(func() {
		f.fields[4].SetText("")
		fieldInfo := activeContest.Contest.GetFieldInfos()
		for _, customField := range activeContest.Contest.ExchangeSent {
			info, ok := fieldInfo[customField]
			if ok && info.FieldType == "tx_const" {
				f.setFieldValue(customField, activeContest.Station.CustomFields[customField])
			} else {
				f.setFieldValue(customField, "")
			}
		}
		for _, customField := range activeContest.Contest.ExchangeRcvd {
			f.setFieldValue(customField, "")
		}
		f.focusField("dx_callsign")
	})

}

func fillToWidth(width int, str string) string {
	for len(str) < width {
		str = str + " "
	}
	return str
}

func leftFillToWidth(width int, str string) string {
	for len(str) < width {
		str = " " + str
	}
	return str
}

type QsoSlice []*gproto.QSO

func (s QsoSlice) Len() int {
	return len(s)
}

func (s QsoSlice) Less(i, j int) bool {
	return s[i].Time < s[j].Time
}

func (s QsoSlice) Swap(i, j int) {
	t := s[i]
	s[i] = s[j]
	s[j] = t
}

func (f *qsoPage) reloadQsoTable() {
	qsos, _ := guiClient.GetActiveQSOs(context.Background(), &emptypb.Empty{})
	if qsos == nil {
		return
	}

	var qsoList QsoSlice = qsos.Qso
	sort.Sort(qsoList)

	f.app.QueueUpdateDraw(func() {
		for i, qso := range qsos.Qso {
			f.qsoTable.SetCell(i+1, 0, tview.NewTableCell(strconv.FormatInt(int64(i+1), 10)).SetAlign(tview.AlignRight))
			f.qsoTable.SetCell(i+1, 1, tview.NewTableCell(formatFreq(qso.Freq)))
			f.qsoTable.SetCell(i+1, 2, tview.NewTableCell(formatDateTime(time.Unix(qso.Time, 0))))
			f.qsoTable.SetCell(i+1, 3, tview.NewTableCell(qso.Mode))
			f.qsoTable.SetCell(i+1, 4, tview.NewTableCell(qso.DxCallsign))
			for j, customField := range activeContest.Contest.ExchangeSent {
				fieldValue, ok := qso.ExchangeSent[customField]
				if ok {
					f.qsoTable.SetCell(i+1, 5+j, tview.NewTableCell(fieldValue))
				}
			}
			for j, customField := range activeContest.Contest.ExchangeRcvd {
				fieldValue, ok := qso.ExchangeRcvd[customField]
				if ok {
					f.qsoTable.SetCell(i+1, 5+len(activeContest.Contest.ExchangeRcvd)+j, tview.NewTableCell(fieldValue))
				}
			}
		}
		f.fields[0].SetText(leftFillToWidth(5, strconv.FormatInt(int64(len(qsos.Qso)+1), 10))).SetFieldStyle(tcell.StyleDefault.Italic(true))
	})
}

func (f *qsoPage) buildQsoTable(app *tview.Application, box *tview.Flex) {
	f.qsoTable = tview.NewTable()
	f.qsoTable.SetFixed(1, 0)
	f.qsoTable.InsertRow(0)
	f.qsoTable.SetBorder(false)
	for i := 0; i <= 10; i++ {
		f.qsoTable.InsertRow(0)
	}
	for i := 0; i < len(f.fields); i++ {
		f.qsoTable.InsertColumn(i)
	}

	titleStyle := tcell.StyleDefault.Bold(true).Foreground(tcell.ColorYellow)

	f.qsoTable.SetCell(0, 0, tview.NewTableCell("").SetStyle(titleStyle))
	f.qsoTable.SetCell(0, 1, tview.NewTableCell("Freq").SetStyle(titleStyle))
	f.qsoTable.SetCell(0, 2, tview.NewTableCell("Date Time").SetStyle(titleStyle))
	f.qsoTable.SetCell(0, 3, tview.NewTableCell("Mode").SetStyle(titleStyle))
	f.qsoTable.SetCell(0, 4, tview.NewTableCell("Call").SetStyle(titleStyle))

	fieldInfo := activeContest.Contest.GetFieldInfos()
	for i, customField := range activeContest.Contest.ExchangeSent {
		info, ok := fieldInfo[customField]
		var fieldName string
		if ok {
			fieldName = info.DisplayName
		}
		f.qsoTable.SetCell(0, 5+i, tview.NewTableCell(fieldName).SetStyle(titleStyle))
	}
	for i, customField := range activeContest.Contest.ExchangeRcvd {
		info, ok := fieldInfo[customField]
		var fieldName string
		if ok {
			fieldName = info.DisplayName
		}
		f.qsoTable.SetCell(0, 5+len(activeContest.Contest.ExchangeSent)+i, tview.NewTableCell(fieldName).SetStyle(titleStyle))
	}
	for i, field := range f.fields {
		f.qsoTable.GetCell(0, i).SetMaxWidth(field.GetFieldWidth())
		f.qsoTable.GetCell(0, i).SetText(fillToWidth(field.GetFieldWidth(), f.qsoTable.GetCell(0, i).Text))
	}
	box.SetDirection(tview.FlexRow)
	box.AddItem(f.qsoTable, 11, 0, false)

	go f.reloadQsoTable()
}

func formatPartialCheckResult(callsign string) string {
	res := &strings.Builder{}

	missing := false
	wrong := false
	for _, ch := range callsign {
		if ch == '!' {
			missing = true
		} else if ch == '@' {
			wrong = true
		} else {
			if missing {
				res.WriteString("[:#660000]")
				res.WriteRune(ch)
				res.WriteString("[-:-:-]")
			} else if wrong {
				res.WriteString("[:#666600]")
				res.WriteRune(ch)
				res.WriteString("[-:-:-]")
			} else {
				res.WriteString("[#00cccc::b]")
				res.WriteRune(ch)
				res.WriteString("[-:-:-]")
			}
			missing = false
			wrong = false
		}
	}

	return res.String()
}

func (f *qsoPage) checkPartial(callsign string) {
	res, err := realTimeGuiClient.CheckPartial(context.Background(), &gproto.CheckPartialRequest{
		Callsign: callsign,
	})

	if err != nil {
		return
	}

	text := &strings.Builder{}
	for _, r := range res.Results {
		fmt.Fprintf(text, `[red::b]%s[-:-:-] `, r.GetTitle())
		for _, c := range r.Callsign {
			text.WriteString(formatPartialCheckResult(c))
			text.WriteByte(' ')
		}
		text.WriteRune('\n')
	}

	f.app.QueueUpdateDraw(func() {
		f.callsignHintList.SetDynamicColors(true)
		f.callsignHintList.SetText(text.String())
	})
}

func (f *qsoPage) buildPartialPanel(app *tview.Application, box *tview.Flex) {
	f.callsignHintList = tview.NewTextView()

	box.AddItem(f.callsignHintList, 0, 1, false)
}

func (f *qsoPage) buildQsoInput(app *tview.Application, box *tview.Flex) {
	f.acceptFunctions = make(map[int]func(textToCheck string, lastChar rune) bool)
	f.focusLossFunctions = make(map[int]func(key tcell.Key))
	f.fieldsNameMap = make(map[string]*tview.InputField)

	qsoIndexField := tview.NewInputField()
	qsoTimeField := tview.NewInputField()
	qsoOtherCallsign := tview.NewInputField()
	qsoFrequencyField := tview.NewInputField()
	qsoModeField := tview.NewInputField()

	f.fields = make([]*tview.InputField, 0)

	qsoIndexField.SetText("(New)").SetFieldWidth(5).SetDisabled(true)
	qsoFrequencyField.SetText("1800.0").SetFieldWidth(8).SetDisabled(true)
	qsoModeField.SetText("SSB").SetFieldWidth(4).SetDisabled(true)

	f.addFieldWithName("", qsoIndexField)
	f.addFieldWithName("freq", qsoFrequencyField)
	f.addFieldWithName("", qsoTimeField)
	f.addFieldWithName("mode", qsoModeField)
	f.addFieldWithName("dx_callsign", qsoOtherCallsign)

	f.acceptFunctions[4] = func(textToCheck string, lastChar rune) bool {
		re := regexp.MustCompile(`^([0-9]+(\.[0-9]*)?)|(([0-9]|[A-Z])?[A-Z]*[0-9]*[A-Z]*)$`)
		return re.MatchString(strings.ToUpper(textToCheck))
	}

	focusNextFieldFunc := func(index int) func(key tcell.Key) {
		return func(key tcell.Key) {
			switch key {
			case tcell.KeyEnter:
				if index == 4 {
					go f.draftQso()
				}
				if index == len(f.fields)-1 {
					go f.submitQso()
				}
				return
			case tcell.KeyTab:
				go func() {
					app.QueueUpdateDraw(
						func() {
							nextField := index + 1
							if nextField == len(f.fields) {
								nextField = 4
							}
							app.SetFocus(f.fields[nextField])
						},
					)
				}()
			}
		}
	}

	qsoOtherCallsign.SetFieldWidth(8).SetFieldBackgroundColor(tcell.NewHexColor(0x333333))

	{
		fieldInfo := activeContest.Contest.GetFieldInfos()
		for _, customField := range activeContest.Contest.ExchangeSent {
			var newField *tview.InputField
			info, ok := fieldInfo[customField]
			if !ok {
				newField = tview.NewInputField()
			} else {
				if info.FieldType == "tx_const" {
					newField = tview.NewInputField()
					newField.SetText(activeContest.Station.CustomFields[customField])
				} else {
					newField = tview.NewInputField()
				}
			}
			newField.SetFieldWidth(8)
			newField.SetFieldBackgroundColor(tcell.NewHexColor(0x333333))
			if newField != nil {
				f.addFieldWithName(customField, newField)
			}
		}
		for _, customField := range activeContest.Contest.ExchangeRcvd {
			var newField *tview.InputField
			_, ok := fieldInfo[customField]
			if !ok {
				newField = tview.NewInputField()
			} else {
				newField = tview.NewInputField()
			}
			newField.SetFieldWidth(8)
			newField.SetFieldBackgroundColor(tcell.NewHexColor(0x333333))
			if newField != nil {
				f.addFieldWithName(customField, newField)
			}
		}
	}

	qsoTimeField.SetFieldWidth(12).SetDisabled(true)
	go clockLoop(app, qsoTimeField)

	qsoOtherCallsign.SetChangedFunc(func(text string) {
		forceUpperCase := strings.ToUpper(text)
		if forceUpperCase != text {
			qsoOtherCallsign.SetText(forceUpperCase)
		}
		go f.checkPartial(forceUpperCase)
	})
	box.SetDirection(tview.FlexRow)

	form := tview.NewFlex()
	form.SetDirection(tview.FlexColumn)
	for i, item := range f.fields {
		form.AddItem(item, item.GetFieldWidth()+1, 0, false)
		f.focusLossFunctions[i] = focusNextFieldFunc(i)
		item.SetAcceptanceFunc(createAcceptFunction(&qsoPageData, i))
		item.SetFinishedFunc(createFocusLossFunction(&qsoPageData, i))
	}

	box.AddItem(form, 1, 0, true)
}

func newModal(p tview.Primitive, width, height int) tview.Primitive {
	return tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(nil, 0, 1, false).
			AddItem(p, height, 1, true).
			AddItem(nil, 0, 1, false), width, 1, true).
		AddItem(nil, 0, 1, false)
}

func (f *qsoPage) handleRadioButton(index int) {
	f.app.QueueUpdateDraw(func() {
		mainPage.ShowPage("radioconfig")
	})
}

func (f *qsoPage) handleRadioModeChange(index int) {
	if f.radioPanel.active {
		//
	}
}

func (f *qsoPage) updateRadioState() {
	radios, err := radioClient.ListRadioStatus(context.Background(), &emptypb.Empty{})
	if err != nil {
		return
	}
	if len(radios.Radios) == 0 {
		f.radioPanel.active = false
	}

	for _, r := range radios.Radios {
		f.radioPanel.active = r.Enabled
		f.app.QueueUpdateDraw(func() {
			f.radioPanel.freq = uint64(r.Rx.Frequency)
			f.radioPanel.freqField.SetText(formatFreq(r.Rx.Frequency))
			f.setFieldValue("freq", formatFreq(r.Rx.Frequency))

			// TODO: ugly
			if r.Rx.Mode == gproto.RadioMode_USB {
				f.radioPanel.mode = "USB"
				f.setFieldValue("mode", "SSB")
				f.radioPanel.modeField.SetCurrentOption(0)
			} else if r.Rx.Mode == gproto.RadioMode_LSB {
				f.radioPanel.mode = "LSB"
				f.setFieldValue("mode", "SSB")
				f.radioPanel.modeField.SetCurrentOption(1)
			} else if r.Rx.Mode == gproto.RadioMode_CW {
				f.radioPanel.mode = "CW"
				f.setFieldValue("mode", "CW")
				f.radioPanel.modeField.SetCurrentOption(2)
			}
		})
		break
	}

	if f.radioPanel.active {
		f.app.QueueUpdateDraw(func() {
			f.radioPanel.configButton.SetLabel("Active")
		})
	} else {
		f.app.QueueUpdateDraw(func() {
			f.radioPanel.configButton.SetLabel("Disabled")
		})
	}
}

func (f *qsoPage) radioStateLoop() {
for_loop:
	for {
		select {
		case <-stopChan:
			break for_loop
		case <-time.After(time.Millisecond * 100):
			f.updateRadioState()
		}
	}
}

func (f *qsoPage) buildRadioPanel(box *tview.Flex) {
	box.SetDirection(tview.FlexColumn)
	f.radioPanel = &radioPanel{}

	f.radioPanel.configButton = tview.NewButton("Disabled")
	f.radioPanel.configButton.SetSelectedFunc(func() { go f.handleRadioButton(0) }).SetBackgroundColor(tcell.ColorRed)
	f.radioPanel.configButton.SetBackgroundColorActivated(tcell.ColorRed)

	f.radioPanel.freqField = tview.NewInputField()
	f.radioPanel.freqField.SetFieldWidth(8)
	f.radioPanel.freqField.SetText("1800.0").SetFieldBackgroundColor(tcell.NewHexColor(0x333333))

	f.radioPanel.modeField = tview.NewDropDown()
	f.radioPanel.modeField.SetOptions([]string{"USB", "LSB", "CW"}, func(text string, index int) {
		f.handleRadioModeChange(0)
	}).SetFieldBackgroundColor(tcell.NewHexColor(0x333333))

	box.AddItem(f.radioPanel.configButton, 10, 0, false)
	box.AddItem(nil, 1, 0, false)
	box.AddItem(f.radioPanel.freqField, 10, 0, false)
	box.AddItem(nil, 1, 0, false)
	box.AddItem(f.radioPanel.modeField, 10, 0, false)

	f.radioPanel.radioDialogue = tview.NewForm()
	f.radioPanel.radioDialogue.SetBorder(true)
	f.radioPanel.radioDialogue.SetTitle("Radio Config")

	f.radioPanel.radioFormModel = tview.NewDropDown().SetLabel("Radio Model")
	f.radioPanel.radioFormUrl = tview.NewInputField().SetLabel("Radio URL")

	for _, v := range supportedRadioList.RadioConfig {
		f.radioPanel.radioFormModel.AddOption(v.Model, nil)
	}

	f.radioPanel.radioDialogue.AddFormItem(f.radioPanel.radioFormModel)
	f.radioPanel.radioDialogue.AddFormItem(f.radioPanel.radioFormUrl)
	f.radioPanel.radioDialogue.AddButton("Confirm", func() {
		go func() {
			_, model := f.radioPanel.radioFormModel.GetCurrentOption()
			radioClient.SetupRadio(context.Background(), &gproto.RadioConfig{
				Channel: "0",
				Model:   model,
				Uri:     f.radioPanel.radioFormUrl.GetText(),
			})

			f.app.QueueUpdateDraw(func() {
				mainPage.SwitchToPage("main")
			})
		}()
	}).AddButton("Cancel", func() {
		mainPage.SwitchToPage("main")
	})

	f.radioPanel.radioDialogue.SetFieldBackgroundColor(tcell.NewHexColor(0x333333))
	mainPage.AddPage("radioconfig", newModal(f.radioPanel.radioDialogue, 60, 20), true, false)

	go f.radioStateLoop()
}

func loggerMain() {
	var err error

	app := tview.NewApplication()
	qsoPageData.app = app

	activeContest, err = guiClient.GetActiveContest(context.Background(), &emptypb.Empty{})
	if err != nil {
		logrus.Fatalf("failed to get active contest: %s", err)
	}

	stopChan = make(chan interface{})

	pages := tview.NewPages()
	mainPage = pages
	mainWindowGrid := tview.NewGrid()
	mainFlex := tview.NewFlex()
	pages.AddPage("main", mainWindowGrid, true, true)

	qsoGrid := tview.NewGrid()

	qsoInputBox := tview.NewBox()
	qsoInput := tview.NewFlex()
	qsoInput.Box = qsoInputBox

	qsoTableBox := tview.NewBox()
	qsoTable := tview.NewFlex()
	qsoTable.Box = qsoTableBox

	callsignCheckerBox := tview.NewBox()
	callsignChecker := tview.NewFlex()
	callsignChecker.Box = callsignCheckerBox
	callsignChecker.Box.SetBorder(true)
	callsignChecker.Box.SetTitle("Checker")

	radioInfoBox := tview.NewBox()
	radioInfo := tview.NewFlex()
	radioInfo.Box = radioInfoBox
	radioInfo.Box.SetBorder(true)
	radioInfo.Box.SetTitle("Radio")

	qsoGrid.SetRows(11, 1)
	qsoGrid.AddItem(qsoTable, 0, 0, 1, 1, 0, 0, false)
	qsoGrid.AddItem(qsoInput, 1, 0, 1, 1, 0, 0, false)
	qsoGrid.AddItem(callsignChecker, 0, 1, 2, 1, 0, 0, false)

	mainFlex.SetDirection(tview.FlexRow)
	mainFlex.AddItem(qsoGrid, 12, 0, false)
	mainFlex.AddItem(radioInfo, 3, 0, true)

	mainWindowGrid.SetRows(-1, 15)
	mainWindowGrid.AddItem(mainFlex, 1, 0, 1, 1, 0, 0, true)

	go func() {
		qsoPageData.buildQsoInput(app, qsoInput)
		qsoPageData.buildQsoTable(app, qsoTable)
		qsoPageData.buildPartialPanel(app, callsignChecker)
		qsoPageData.buildRadioPanel(radioInfo)

		qsoInputWidth := 0
		for _, field := range qsoPageData.fields {
			qsoInputWidth += field.GetFieldWidth()
		}
		qsoInputWidth += len(qsoPageData.fields)
		qsoGrid.SetColumns(qsoInputWidth, -1)

		if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
			panic(err)
		}

		close(stopChan)
	}()

	<-stopChan
}
