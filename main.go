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
	if err := os.MkdirAll(databaseDir, os.ModePerm); err != nil {
		logrus.Fatalf("Failed to create directory: %v", err)
	}
	systemLogPath := path.Join(databaseDir, "systemlog.log")
	if fp, err := os.OpenFile(systemLogPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm); err != nil {
		logrus.Fatal("Cannot open path for log: ", err)
	} else {
		logrus.SetOutput(fp)
	}
	databaseName := path.Join(databaseDir, "log.sqlite3")
	if err := logdb.Init(databaseName); err != nil {
		logrus.Fatalf("Failed to setup main database (%s): %v", databaseName, err)
	}

	// Init global config
	state.LoadState(path.Join(databaseDir, "config.json"))
}

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)
	app, err := gtk.ApplicationNew("dev.matsu.contestlog", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		logrus.Fatal("Failed to init gtk application: ", err)
	}

	matsuLogInit()

	builder, err := gtk.BuilderNewFromString(resources.Glade)
	if err != nil {
		logrus.Fatalf("Unable to init builder: %v", err)
	}
	state.GetState().Gui = builder

	configWindow.Init(builder)
	InitMainWindow(builder, app)
	InitRigctrl()
	go webui.SetupServer()

	gtk.Main()

	state.SaveState()
}
