package service

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/dvordrova/gop/tui/service/handlers"
)

type Layout interface {
	Update(msg tea.Msg) tea.Cmd
	Name() string
	View() string
	IsOk() bool
	Keys() help.KeyMap
	IsEditable() bool
	SetNotEditable()
	SetEditable()
}

type Model struct {
	Layouts          []Layout
	currentLayoutInd int
	fullScreen       bool
	KeyMap           KeyMap
	help             help.Model
}

func NewModel() Model {
	return Model{
		Layouts: []Layout{
			handlers.NewModel(),
			handlers.NewModel(),
			handlers.NewModel(),
			handlers.NewModel(),
			handlers.NewModel(),
		},
		currentLayoutInd: 0,
		KeyMap:           DefaultKeyMap,
		fullScreen:       true,
		help:             help.New(),
	}
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (s Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	wasDefault := false
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, s.KeyMap.Quit):
			return s, tea.Quit
		case key.Matches(msg, s.KeyMap.FullScreen):
			s.fullScreen = !s.fullScreen
			return s, altScreen[s.fullScreen]
		case key.Matches(msg, s.KeyMap.Help):
			s.help.ShowAll = !s.help.ShowAll
		case key.Matches(msg, s.KeyMap.Left) && !s.Layouts[s.currentLayoutInd].IsEditable():
			s.currentLayoutInd = (s.currentLayoutInd - 1 + len(s.Layouts)) % len(s.Layouts)
			s.Layouts[s.currentLayoutInd].SetNotEditable()
		case key.Matches(msg, s.KeyMap.PrevTab):
			s.currentLayoutInd = (s.currentLayoutInd - 1 + len(s.Layouts)) % len(s.Layouts)
			s.Layouts[s.currentLayoutInd].SetEditable()
		case key.Matches(msg, s.KeyMap.Right) &&
			!s.Layouts[s.currentLayoutInd].IsEditable():
			s.currentLayoutInd = (s.currentLayoutInd + 1) % len(s.Layouts)
			s.Layouts[s.currentLayoutInd].SetNotEditable()
		case key.Matches(msg, s.KeyMap.NextTab):
			s.currentLayoutInd = (s.currentLayoutInd + 1) % len(s.Layouts)
			s.Layouts[s.currentLayoutInd].SetEditable()
		default:
			wasDefault = true
		}
	}

	if wasDefault {
		return s, s.Layouts[s.currentLayoutInd].Update(msg)
	} else {
		return s, nil
	}
}

func tabBorderWithBottom(left, middle, right string) lipgloss.Border {
	border := lipgloss.RoundedBorder()
	border.BottomLeft = left
	border.Bottom = middle
	border.BottomRight = right
	return border
}

var altScreen = map[bool]tea.Cmd{
	true:  tea.EnterAltScreen,
	false: tea.ExitAltScreen,
}
