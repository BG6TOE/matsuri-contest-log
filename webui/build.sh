#!/bin/bash

set -xe

SOURCE_DIR=$(readlink -f $(dirname "${BASH_SOURCE[0]}"))

generate_rpc() {
    cd $SOURCE_DIR
    $HOME/go/bin/webrpc-gen -schema ./rpcserver/rpc.ridl -target js -pkg=rpc_gen -client -out web/rpc.gen.js
    $HOME/go/bin/webrpc-gen -schema ./rpcserver/rpc.ridl -target go -pkg=rpc_gen -server -out rpc_gen/rpc.gen.go
}

generate_ui() {
    cd $SOURCE_DIR/ui-project
    yarn build
    mkdir -p ../web-dist/
    cp -r dist/* ../web-dist/
}

generate_rpc
generate_ui