package webui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"matsu.dev/matsuri-contest-log/webui/rpc_gen"
	"matsu.dev/matsuri-contest-log/webui/rpcserver"
	"matsu.dev/matsuri-contest-log/webui/ws"
)

//go:embed web-dist/*
var webResources embed.FS

func SetupServer() {
	r := mux.NewRouter()
	mclService := rpc_gen.NewMCLServiceServer(&rpcserver.MCLServer{})

	httpFs, err := fs.Sub(webResources, "web-dist")
	if err != nil {
		logrus.Fatalf("Failed to init webui files: %v", err)
	}

	r.PathPrefix("/rpc/MCLService/").Handler(mclService)
	r.PathPrefix("/web/").Handler(http.StripPrefix("/web/", http.FileServer(http.FS(httpFs))))
	r.Path("/ws").HandlerFunc(ws.HandleWebsocketConnection)

	logrus.Fatal(http.ListenAndServe(":22222", r))
}
