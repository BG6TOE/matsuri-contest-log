package main

import _ "embed"

var (
	//go:embed gui.css
	cssData string

	//go:embed gui.glade
	guiDesc string
)
