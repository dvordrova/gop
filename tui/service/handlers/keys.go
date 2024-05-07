package handlers

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up                key.Binding
	Down              key.Binding
	PrevField         key.Binding
	NextField         key.Binding
	Delete            key.Binding
	AutoComplete      key.Binding
	CharacterBackward key.Binding
	CharacterForward  key.Binding
}

var DefaultKeyMap = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up"),
		key.WithHelp("↑", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down"),
		key.WithHelp("↓", "move down"),
	),
	PrevField: key.NewBinding(
		key.WithKeys("left"),
		key.WithHelp("<-", "move left"),
	),
	NextField: key.NewBinding(
		key.WithKeys("right", " ", "enter"),
		key.WithHelp("->/space/enter", "move right"),
	),
	Delete: key.NewBinding(
		key.WithKeys("ctrl+backspace"),
		key.WithHelp("ctrl+backspace", "delete row"),
	),
	AutoComplete: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "autocomplete"),
	),
	CharacterBackward: key.NewBinding(
		key.WithKeys("<"),
		key.WithHelp("<", "move left"),
	),
	CharacterForward: key.NewBinding(
		key.WithKeys(">"),
		key.WithHelp(">", "move right"),
	),
}

func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.CharacterBackward, k.CharacterForward}
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.CharacterForward, k.CharacterForward, k.Up, k.Down, k.PrevField, k.NextField},
		{k.Delete},
	}
}
