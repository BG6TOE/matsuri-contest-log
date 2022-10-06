package main

import (
	"flag"

	"github.com/sirupsen/logrus"
	embed "matsu.dev/matsuri-contest-log/embed"
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

	embed.Start(&embed.EmbedConfig{RPCHost: "tcp://127.0.0.1:62122"})

	select {}
}
