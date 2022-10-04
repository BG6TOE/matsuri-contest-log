package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/url"
	"os"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	binlog "matsu.dev/matsuri-contest-log/binlogserver"
	gui "matsu.dev/matsuri-contest-log/guiserver"
	"matsu.dev/matsuri-contest-log/version"
)

var (
	database      = flag.String("database", "contest.sqlite", "The path to the QSO database of current contest.")
	contestConfig = flag.String("contest-config", "", "The config for creating a new contest database if the database does not exist.")
	rpcHost       = flag.String("rpc-host", "tcp://0.0.0.0:62122", "The address of the gRPC server.")
)

func main() {
	flag.Parse()

	logrus.Infof("MCL Version %s", version.Version)
	logrus.Infof("Build time: %s", version.BuildTime)
	logrus.Infof("Commit: %s", version.GitCommit)

	if info, err := os.Stat(*database); os.IsNotExist(err) {
		manifest := binlog.ContextManifest{}
		json.Unmarshal([]byte(*contestConfig), &manifest)
		contestUuid, err := uuid.NewUUID()
		if err != nil {
			logrus.Panic(err)
		}
		manifest.Uid = contestUuid.String()
		manifest.Filename = *database
		binlog.NewContest(manifest)
	} else if err != nil {
		logrus.Panic(err)
	} else if info.IsDir() {
		logrus.Panic("%s is a dir", *database)
	}

	url, err := url.Parse(*rpcHost)
	if err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
	lis, err := net.Listen(url.Scheme, url.Host)
	if err != nil {
		panic(fmt.Errorf("failed to create server: %v", err))
	}
	logrus.Infof("Listening at %s", lis.Addr().String())
	grpcServer := grpc.NewServer()

	binlogServer := binlog.NewServer(&binlog.BinlogServerConfig{GrpcServer: grpcServer, Database: *database})
	guiServer := gui.NewServer(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Panicf("failed to serve: %v", err)
		}
	}()

	binlogServer.Init(&binlog.BinlogServerConfig{Database: *database})
	err = guiServer.Init(&gui.GuiServerConfig{BinlogServerAddr: *rpcHost})
	if err != nil {
		logrus.Panicf("failed to init gui server: %v", err)
	}

	select {}
}
