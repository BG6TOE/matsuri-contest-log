package main

import (
	"fmt"
	"math/rand"
	"os"

	"github.com/rivo/tview"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"matsu.dev/matsuri-contest-log/embed"
	gproto "matsu.dev/matsuri-contest-log/proto"
)

var (
	guiClient         gproto.GuiClient
	realTimeGuiClient gproto.RealtimeGuiClient

	selectedDatabase string
)

func createContest() {
	app := tview.NewApplication()
	flex := tview.NewFlex()
	formContest := tview.NewForm().
		AddInputField("Contest Definition", "", 20, nil, nil).
		AddInputField("Contest Name", "new contest", 30, nil, nil).
		AddInputField("Contest Start", "", 20, nil, nil).
		AddInputField("Contest End", "", 20, nil, nil)
	formStation := tview.NewForm().
		AddInputField("Callsign", "", 30, nil, nil).
		AddInputField("Grid", "", 8, nil, nil).
		AddDropDown("ITU Region", []string{"1", "2", "3"}, 4, nil)
	formAction := tview.NewForm().
		AddInputField("Database", "", 30, nil, nil).
		AddButton("Create", nil).
		AddButton("Quit", func() {
			app.Stop()
		})
	flex.AddItem(formContest, 0, 1, true)
	flex.AddItem(formStation, 0, 1, false)
	flex.AddItem(formAction, 0, 1, false)
	flex.SetBorder(true).SetTitle("Create Contest").SetTitleAlign(tview.AlignLeft)
	if err := app.SetRoot(flex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func guiMain() {
	box := tview.NewBox()
	if err := tview.NewApplication().SetRoot(box, true).Run(); err != nil {
		panic(err)
	}
}

func main() {
	randomSocketName := make([]byte, 8)
	rand.Read(randomSocketName)
	serverAddress := fmt.Sprintf("unix:///tmp/mcl.%x.socket", randomSocketName)

	embed.Run(&embed.EmbedConfig{RPCHost: serverAddress})

	conn, err := grpc.Dial(serverAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("Failed to connect to the server: %v", err)
	}
	guiClient = gproto.NewGuiClient(conn)
	realTimeGuiClient = gproto.NewRealtimeGuiClient(conn)

	if len(os.Args) == 1 {
		logrus.Infof("no database opened")
		createContest()
	}

	guiMain()
}
