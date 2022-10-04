package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
	"github.com/pkg/errors"
	"matsu.dev/matsuri-contest-log/embed"
)

// vmArguments may be set by hover at compile-time
var vmArguments string

func main() {
	// DO NOT EDIT, add options in options.go
	mainOptions := []flutter.Option{
		flutter.OptionVMArguments(strings.Split(vmArguments, ";")),
		flutter.WindowIcon(iconProvider),
	}

	err := flutter.Run(append(options, mainOptions...)...)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//  Make sure to use the same channel name as was used on the Flutter client side.
const mclEmbedChannelName = "mcl.matsu.dev/embed"

// MyBatteryPlugin demonstrates how to call a platform-specific API to retrieve
// the current battery level.
//
// This example matches the guide example available on:
// https://flutter.dev/docs/development/platform-integration/platform-channels
type MclEmbedPlugin struct{}

var _ flutter.Plugin = &MclEmbedPlugin{} // compile-time type check

// InitPlugin creates a MethodChannel and set a HandleFunc to the
// shared 'getBatteryLevel' method.
// https://godoc.org/github.com/go-flutter-desktop/go-flutter/plugin#MethodChannel
//
// The plugin is using a MethodChannel through the StandardMethodCodec.
//
// You can also use the more basic BasicMessageChannel, which supports basic,
// asynchronous message passing using a custom message codec.
// You can also use the specialized BinaryCodec, StringCodec, and JSONMessageCodec
// struct, or create your own codec.
//
// The JSONMessageCodec was deliberately left out because in Go its better to
// decode directly to known structs.
func (p *MclEmbedPlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, mclEmbedChannelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("start", handleMclEmbedStart)
	return nil // no error
}

// handleGetBatteryLevel is called when the method getBatteryLevel is invoked by
// the dart code.
// The function returns a fake battery level, without raising an error.
//
// Supported return types of StandardMethodCodec codec are described in a table:
// https://godoc.org/github.com/go-flutter-desktop/go-flutter/plugin#StandardMessageCodec
func handleMclEmbedStart(database interface{}) (reply interface{}, err error) {
	embed.Start(&embed.EmbedConfig{
		Database: database.(string),
		RPCHost:  "tcp://127.0.0.1:62122",
	})
	return nil, nil
}

func iconProvider() ([]image.Image, error) {
	execPath, err := os.Executable()
	if err != nil {
		return nil, errors.Wrap(err, "failed to resolve executable path")
	}
	execPath, err = filepath.EvalSymlinks(execPath)
	if err != nil {
		return nil, errors.Wrap(err, "failed to eval symlinks for executable path")
	}
	imgFile, err := os.Open(filepath.Join(filepath.Dir(execPath), "assets", "icon.png"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open assets/icon.png")
	}
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode image")
	}
	return []image.Image{img}, nil
}
