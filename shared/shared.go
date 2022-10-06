package main

// #include <stdlib.h>

import (
	"C"
)
import "matsu.dev/matsuri-contest-log/embed"

//export run_server
func run_server(rpcHost *C.char) C.int {

	embed.Start(&embed.EmbedConfig{RPCHost: C.GoString(rpcHost)})

	return 0
}

func main() {}
