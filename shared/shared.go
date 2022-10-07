package main

// #include <stdlib.h>

import (
	"C"
)
import "matsu.dev/matsuri-contest-log/embed"

//export run_mcl_gui_server
func run_mcl_gui_server(rpcHost *C.char) C.int {

	embed.Run(&embed.EmbedConfig{RPCHost: C.GoString(rpcHost)})

	return 0
}

func main() {}
