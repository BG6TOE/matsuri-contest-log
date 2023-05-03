package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"matsu.dev/matsuri-contest-log/embed"
	gproto "matsu.dev/matsuri-contest-log/proto"

	"github.com/gin-gonic/gin"
)

var (
	supportedRadioList *gproto.SupportedRadioList

	selectedDatabase string
)

func main() {
	randomSocketName := make([]byte, 8)
	rand.Read(randomSocketName)
	serverAddress := embed.RunWithLocalSocket(false)
	staticFiles := os.Getenv("STATIC")

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

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.POST("/call/:class/:api", handleHttpApiCall)
	r.POST("/support/filepicker", filePicker)

	if os.Getenv("FRONTEND_SERVER") != "" {
		r.Any("/web/*proxyPath", proxy)
	} else {
		r.Static("/web", staticFiles)
	}

	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		logrus.Fatalf("Failed to start internal http server: %v", err)
	}
	logrus.Infof("Listening at: tcp://%s", listener.Addr().String())
	logrus.Infof("Frontend available at: http://%s/web", listener.Addr().String())
	logrus.Fatalf("Failed to start internal http server: %v", http.Serve(listener, r.Handler()))
}
