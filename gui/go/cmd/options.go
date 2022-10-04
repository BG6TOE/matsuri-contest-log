package main

import (
	"github.com/go-flutter-desktop/go-flutter"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(600, 400),
	flutter.WindowDimensionLimits(600, 400, 100000, 100000),
	flutter.AddPlugin(&MclEmbedPlugin{}),
}
