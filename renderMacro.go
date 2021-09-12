package main

import (
	"strings"
	"text/template"

	"matsu.dev/matsuri-contest-log/state"
)

type macroKeyRenderer struct {
}

type MacroRenderContext struct {
	MyCall  string
	ExchRST string
	Exch    string
	DxCall  string
}

var macroRenderer macroKeyRenderer

func (r *macroKeyRenderer) RenderMacroKey(keyPressed string, ctx *MacroRenderContext) string {
	t := ""
	for _, v := range state.GetState().MacroKeyMap {
		if v.Key == keyPressed {
			t = v.Value
			break
		}
	}
	if t == "" {
		return ""
	}

	output := &strings.Builder{}
	tmpl, err := template.New(keyPressed).Parse(t)
	if err != nil {
		return ""
	}
	if err := tmpl.Execute(output, ctx); err != nil {
		return ""
	}

	return output.String()
}
