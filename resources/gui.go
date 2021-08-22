package resources

import _ "embed"

var (
	//go:embed gui.css
	CSS string

	//go:embed gui.glade
	Glade string

	//go:embed dbinit.sql
	DatabaseStructureSQL string
)
