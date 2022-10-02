package main

import (
	"encoding/json"
	"flag"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"matsu.dev/matsuri-contest-log/server"
)

var (
	database      = flag.String("database", "contest.sqlite", "The path to the QSO database of current contest.")
	contestConfig = flag.String("contest-config", "", "The config for creating a new contest database if the database does not exist.")
	rpcHost       = flag.String("rpc-host", "tcp://0.0.0.0:62122", "The address of the gRPC server.")
)

func main() {
	flag.Parse()

	logrus.Infof("MCL Version %s", Version)
	logrus.Infof("Build time: %s", BuildTime)
	logrus.Infof("Commit: %s", GitCommit)

	if info, err := os.Stat(*database); os.IsNotExist(err) {
		manifest := server.ContextManifest{}
		json.Unmarshal([]byte(*contestConfig), &manifest)
		contestUuid, err := uuid.NewUUID()
		if err != nil {
			logrus.Panic(err)
		}
		manifest.Uid = contestUuid.String()
		manifest.Filename = *database
		server.NewContest(manifest)
	} else if err != nil {
		logrus.Panic(err)
	} else if info.IsDir() {
		logrus.Panic("%s is a dir", *database)
	}

	serv := server.NewServer(*database, *rpcHost)

	logrus.Infof("Started server: %p", serv)

	select {}
}
