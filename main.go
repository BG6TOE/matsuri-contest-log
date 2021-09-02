package main

import (
	"os"
	"path"

	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/sirupsen/logrus"

	logdb "matsu.dev/matsuri-contest-log/logdb"
	resources "matsu.dev/matsuri-contest-log/resources"
	state "matsu.dev/matsuri-contest-log/state"
	"matsu.dev/matsuri-contest-log/webui"
)

func matsuLogInit() {
	// Init log db
	homeDir := glib.GetHomeDir()
	databaseDir := path.Join(homeDir, ".config", "dev.matsu.contestlog")
	databaseName := path.Join(databaseDir, "log.sqlite3")
	if err := os.MkdirAll(databaseDir, os.ModeDir); err != nil {
		logrus.Fatalf("Failed to create directory: %v", err)
	}
	if err := logdb.Init(databaseName); err != nil {
		logrus.Fatalf("Failed to setup main database: %v", err)
	}

	// Init global config
	state.LoadState(path.Join(databaseDir, "config.json"))
}

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	matsuLogInit()

	builder, err := gtk.BuilderNewFromString(resources.Glade)
	if err != nil {
		logrus.Fatalf("Unable to init builder: %v", err)
	}
	state.GetState().Gui = builder

	InitConfigDialog(builder)
	InitMainWindow(builder)
	InitRigctrl()
	go webui.SetupServer()

	gtk.Main()

	state.SaveState(path.Join(glib.GetHomeDir(), ".config", "dev.matsu.contestlog", "config.json"))
}
