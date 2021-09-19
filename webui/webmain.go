package webui

import (
	"embed"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"matsu.dev/matsuri-contest-log/webui/rpc_gen"
	"matsu.dev/matsuri-contest-log/webui/rpcserver"
	"matsu.dev/matsuri-contest-log/webui/ws"
)

//go:embed web/*
var webResources embed.FS

func SetupServer() {
	r := mux.NewRouter()
	mclService := rpc_gen.NewMCLServiceServer(&rpcserver.MCLServer{})
	r.PathPrefix("/rpc/MCLService/").Handler(mclService)
	r.PathPrefix("/web/").Handler(http.FileServer(http.FS(webResources)))
	r.PathPrefix("/web-dev/").Handler(http.StripPrefix("/web-dev/", http.FileServer(http.Dir("./webui/web/"))))
	r.Path("/ws").HandlerFunc(ws.HandleWebsocketConnection)

	logrus.Fatal(http.ListenAndServe(":22222", r))
}
