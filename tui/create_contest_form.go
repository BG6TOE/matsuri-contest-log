package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	gproto "matsu.dev/matsuri-contest-log/proto"
)

func createContest() {
	app := tview.NewApplication()
	flex := tview.NewFlex()

	var contestDescriptor string
	var contestDefine *gproto.ContestMetadata

	var (
		formContest *tview.Form
		formAction  *tview.Form
	)

	var (
		contestNameField *tview.InputField
		contestStartTime *tview.InputField
		contestEndTime   *tview.InputField
		callsignField    *tview.InputField
	)

	flex.SetDirection(tview.FlexRow)

	contestStartTime = tview.NewInputField().SetFieldWidth(16).SetLabel("Contest Start (UTC)").SetPlaceholder("YYYY-MM-DD HH:MM").SetChangedFunc(func(text string) {
		re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}$`)
		if re.Match([]byte(text)) {
			contestStartTime.SetFieldTextColor(tcell.ColorWhite)
		} else {
			contestStartTime.SetFieldTextColor(tcell.ColorRed)
		}
	})
	contestEndTime = tview.NewInputField().SetFieldWidth(16).SetLabel("Contest End (UTC)").SetPlaceholder("YYYY-MM-DD HH:MM").SetChangedFunc(func(text string) {
		re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2} \d{2}:\d{2}$`)
		if re.Match([]byte(text)) {
			contestEndTime.SetFieldTextColor(tcell.ColorWhite)
		} else {
			contestEndTime.SetFieldTextColor(tcell.ColorRed)
		}
	})
	callsignField = tview.NewInputField().SetFieldWidth(16).SetLabel("Callsign").SetChangedFunc(func(text string) {
		forceUpperCase := strings.ToUpper(text)
		if forceUpperCase != text {
			callsignField.SetText(forceUpperCase)
		}
	})
	contestNameField = tview.NewInputField().SetFieldWidth(16).SetLabel("Contest Name")

	formContest = tview.NewForm().
		AddInputField("Contest Definition", "", 20, nil, func(text string) {
			file, err := os.Open(text)
			if err != nil {
				return
			}
			defer file.Close()
			data, err := ioutil.ReadAll(file)
			if err != nil {
				return
			}
			contestDescriptor = string(data)

			contest, err := guiClient.ParseContest(context.Background(), &gproto.ParseContestRequest{
				ContestDescriptor: contestDescriptor,
			})
			if err != nil {
				return
			}
			for formContest.GetFormItemCount() > 5 {
				formContest.RemoveFormItem(formContest.GetFormItemCount() - 1)
			}
			contestDefine = contest
			fieldInfo := contestDefine.GetFieldInfos()
			for _, customField := range contestDefine.ExchangeSent {
				info, ok := fieldInfo[customField]
				if !ok {
					formContest.AddInputField(customField, "", 0, nil, nil)
				} else {
					if info.FieldType == "tx_const" {
						formContest.AddInputField(info.DisplayName, "", 0, nil, nil)
					}
				}
			}
		}).
		AddFormItem(contestNameField).
		AddFormItem(contestStartTime).
		AddFormItem(contestEndTime).
		AddFormItem(callsignField)

	parseDateTime := func(str string) time.Time {
		var (
			y, m, d, H, M int
		)
		fmt.Sscanf(str, "%d-%d-%d %d:%d", y, m, d, H, M)

		return time.Date(y, time.Month(m), d, H, M, 0, 0, time.UTC)
	}

	createContest := func() {
		customFields := make(map[string]string)
		fieldInfo := contestDefine.GetFieldInfos()
		fieldIndex := 0
		for _, customField := range contestDefine.ExchangeSent {
			info, ok := fieldInfo[customField]
			if !ok {
				customFields[customField] = formContest.GetFormItem(fieldIndex + 5).(*tview.InputField).GetText()
				fieldIndex++
			} else {
				if info.FieldType == "tx_const" {
					customFields[customField] = formContest.GetFormItem(fieldIndex + 5).(*tview.InputField).GetText()
					fieldIndex++
				}
			}
		}
		req := &gproto.CreateContestRequest{
			DatabaseName: formAction.GetFormItem(0).(*tview.InputField).GetText(),
			Contest: &gproto.ActiveContest{
				Name:    contestNameField.GetText(),
				Contest: contestDefine,
				Station: &gproto.Station{
					Callsign:     callsignField.GetText(),
					Grid:         "OO00",
					Region:       0,
					CustomFields: customFields,
				},
				BeginTimestamp: parseDateTime(contestStartTime.GetText()).Unix(),
				EndTimestamp:   parseDateTime(contestEndTime.GetText()).Unix(),
				ContestScript:  contestDescriptor,
			},
		}
		resp, err := guiClient.CreateContest(context.Background(), req)
		if err != nil {
			app.Stop()
			logrus.Fatalf("failed to create contest: %v", err)
		}
		if resp.ResultCode != gproto.ResultCode_success {
			app.Stop()
			logrus.Fatalf("failed to create contest: %s (err=%d)", resp.ErrorMessage, resp.ResultCode)
		}
		selectedDatabase = formAction.GetFormItem(0).(*tview.InputField).GetText()
		app.Stop()
	}

	formAction = tview.NewForm().
		AddInputField("Database", "", 20, nil, nil).
		AddButton("Create", createContest).
		AddButton("Quit", func() {
			app.Stop()
		})

	flex.AddItem(formContest, 20, 1, true)
	flex.AddItem(formAction, 0, 1, false)
	flex.SetBorder(true).SetTitle("Create Contest").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
