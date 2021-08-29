#!/bin/bash

OS="$(uname)"

echo "Current OS is $OS"

build_Linux() {
    output="out/linux/$(uname -m)"
    mkdir -p "$output"
    go build -o "$output"/MatsuriLog
}

build_Mingw() {
    output="out/windows/$(uname -m)"
    mkdir -p $output
    go build -o "$output"/MatsuriLog.exe -ldflags -H=windowsgui
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
