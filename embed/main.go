package embed

import (
	"encoding/json"
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

type EmbedConfig struct {
	Database      string
	ContestConfig string
	RPCHost       string
}

func Start(conf *EmbedConfig) {
	logrus.Infof("MCL Version %s", version.Version)
	logrus.Infof("Build time: %s", version.BuildTime)
	logrus.Infof("Commit: %s", version.GitCommit)

	if info, err := os.Stat(conf.Database); os.IsNotExist(err) {
		manifest := binlog.ContextManifest{}
		json.Unmarshal([]byte(conf.ContestConfig), &manifest)
		contestUuid, err := uuid.NewUUID()
		if err != nil {
			logrus.Panic(err)
		}
		manifest.Uid = contestUuid.String()
		manifest.Filename = conf.Database
		binlog.NewContest(manifest)
	} else if err != nil {
		logrus.Panic(err)
	} else if info.IsDir() {
		logrus.Panic("%s is a dir", conf.Database)
	}

	url, err := url.Parse(conf.RPCHost)
	if err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
	lis, err := net.Listen(url.Scheme, url.Host)
	if err != nil {
		panic(fmt.Errorf("failed to create server: %v", err))
	}
	logrus.Infof("Listening at %s", lis.Addr().String())
	grpcServer := grpc.NewServer()

	binlogServer := binlog.NewServer(&binlog.BinlogServerConfig{GrpcServer: grpcServer, Database: conf.Database})
	guiServer := gui.NewServer(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Panicf("failed to serve: %v", err)
		}
	}()

	binlogServer.Init(&binlog.BinlogServerConfig{Database: conf.Database})
	err = guiServer.Init(&gui.GuiServerConfig{BinlogServerAddr: conf.RPCHost})
	if err != nil {
		logrus.Panicf("failed to init gui server: %v", err)
	}
}
