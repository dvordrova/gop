package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
)

type HelpTogether struct {
	keyMaps []help.KeyMap
}

func (k HelpTogether) ShortHelp() []key.Binding {
	res := []key.Binding{}
	for _, km := range k.keyMaps {
		res = append(res, km.ShortHelp()...)
	}

	return res
}

func (k HelpTogether) FullHelp() [][]key.Binding {
	res := make([][]key.Binding, 0, len(k.keyMaps))
	for _, km := range k.keyMaps {
		res = append(res, km.FullHelp()...)
	}
	return res
}

func HelpView(h help.Model, keyMaps ...help.KeyMap) string {
	return h.View(HelpTogether{keyMaps: keyMaps})
}
