package embed

import (
	"fmt"
	"net"
	"net/url"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	binlog "matsu.dev/matsuri-contest-log/binlogserver"
	gui "matsu.dev/matsuri-contest-log/guiserver"
	"matsu.dev/matsuri-contest-log/radio"
	"matsu.dev/matsuri-contest-log/version"
)

type EmbedConfig struct {
	RPCHost string
}

func Run(conf *EmbedConfig) {
	logrus.Infof("MCL Version %s", version.Version)
	logrus.Infof("Build time: %s", version.BuildTime)
	logrus.Infof("Commit: %s", version.GitCommit)

	url, err := url.Parse(conf.RPCHost)
	if err != nil {
		panic(fmt.Errorf("failed to start server: %v", err))
	}
	lis, err := net.Listen(url.Scheme, url.Host)
	if err != nil {
		panic(fmt.Errorf("failed to create server: %v", err))
	}
	logrus.Infof("listening at %s", lis.Addr().String())
	grpcServer := grpc.NewServer()

	binlog.NewRpcServer(grpcServer)
	radio.NewServer(grpcServer)
	guiServer := gui.NewServer(grpcServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			logrus.Panicf("failed to serve: %v", err)
		}
	}()

	err = guiServer.Init(&gui.GuiServerConfig{BinlogServerAddr: conf.RPCHost})
	if err != nil {
		logrus.Panicf("failed to init gui server: %v", err)
	}
}
