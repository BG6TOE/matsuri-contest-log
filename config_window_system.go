package main

import (
	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

func (c *ConfigWindow) initSystemConfig(builder *gtk.Builder) {
	mustGetObj(builder, "config-version-label").(*gtk.Label).SetText(fmt.Sprintf("Version: %s\nBuild Time: %s", GitCommit, BuildTime))
}
