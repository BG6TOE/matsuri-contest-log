package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/url"
	"os"
	"syscall"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"matsu.dev/matsuri-contest-log/embed"
	gproto "matsu.dev/matsuri-contest-log/proto"
)

var (
	guiClient         gproto.GuiClient
	radioClient       gproto.RadioClient
	realTimeGuiClient gproto.RealtimeGuiClient

	supportedRadioList *gproto.SupportedRadioList

	selectedDatabase string
)

func main() {
	randomSocketName := make([]byte, 8)
	rand.Read(randomSocketName)
	serverAddress := "tcp://127.5.9.73:5973"

	fp, err := os.OpenFile(fmt.Sprintf("/tmp/mcl.%x.log", randomSocketName), os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		logrus.Fatalf("failed to open file for log: %v", err)
	}
	defer fp.Close()
	logrus.SetOutput(fp)

	embed.Run(&embed.EmbedConfig{RPCHost: serverAddress})

	syscall.Unlink(serverAddress)

	conn, err := grpc.Dial(serverAddress, grpc.WithContextDialer(func(ctx context.Context, s string) (net.Conn, error) {
		var d net.Dialer
		url, err := url.Parse(serverAddress)
		if err != nil {
			panic(fmt.Errorf("failed to start server: %v", err))
		}
		if url.Scheme == "unix" {
			url.Host = url.Path
		}
		return d.DialContext(ctx, url.Scheme, url.Host)
	}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatalf("Failed to connect to the server: %v", err)
	}
	guiClient = gproto.NewGuiClient(conn)
	realTimeGuiClient = gproto.NewRealtimeGuiClient(conn)
	radioClient = gproto.NewRadioClient(conn)

	if len(os.Args) == 1 {
		logrus.Infof("no database opened")
		createContest()
	} else {
		selectedDatabase = os.Args[1]
	}

	guiClient.LoadContest(context.Background(), &gproto.LoadContestRequest{
		DatabaseName: selectedDatabase,
	})

	supportedRadioList, err = radioClient.ListSupportedRadios(context.Background(), &emptypb.Empty{})
	if err != nil {
		logrus.Errorf("Failed to load supported radio list: %v", err)
	}

	loggerMain()
}
