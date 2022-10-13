#!/bin/bash

set -xe

OS="$(uname)"

COMMITID=$(git rev-parse HEAD)
BUILDTIME=$(date -u '+%Y-%m-%dT%H:%M:%SZ')
VERSION=$(git describe --tags)

GOLDFLAGS="-X version.GitCommit=$COMMITID -X version.BuildTime=$BUILDTIME -X version.Version=$VERSION"

echo "Current OS is $OS"

build_Linux() {
    output="out/linux/$(uname -m)"
    mkdir -p "$output"
    go build -o "$output"/MatsuriLog -ldflags "$GOLDFLAGS"
}

build_Mingw() {
    output="out/windows/$(uname -m)"
    mkdir -p $output
    go build -o "$output"/MatsuriLog.exe -ldflags "$GOLDFLAGS -H=windowsgui"
    cd $output
    ldd MatsuriLog.exe | grep '\/mingw.*\.dll' -o | xargs -I{} cp -n "{}" .
}

if uname | grep -q "Linux" -; then
 build_Linux
 exit 0
fi

if uname | grep -q "MINGW" -; then
 build_Mingw
 exit 0
fi

echo "Unsupported OS!"
exit 1
