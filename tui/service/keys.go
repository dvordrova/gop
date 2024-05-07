package service

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Help       key.Binding
	Quit       key.Binding
	FullScreen key.Binding
	PrevTab    key.Binding
	Left       key.Binding
	NextTab    key.Binding
	Right      key.Binding
}

var DefaultKeyMap = KeyMap{
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("esc", "ctrl+c"),
		key.WithHelp("esc/ctrl+c", "quit"),
	),
	FullScreen: key.NewBinding(
		key.WithKeys("ctrl+f"),
		key.WithHelp("ctrl+f", "toggle full screen mode"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("[", "{", "х", "Х"),
		key.WithHelp("[", "prev tab"),
	),
	Left: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("<-", "move left"),
	),
	NextTab: key.NewBinding(
		key.WithKeys("]", "}", "ъ", "Ъ"),
		key.WithHelp("]", "next tab"),
	),
	Right: key.NewBinding(
		key.WithKeys("right"),
		key.WithHelp("->", "move right"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.PrevTab, k.NextTab, k.FullScreen, k.Help, k.Quit}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.PrevTab, k.NextTab},
		{k.FullScreen, k.Help, k.Quit},
	}
}
