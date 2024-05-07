package handlers

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Handlers        []*HandlerRow
	ChosenHandler   int
	ChosenAttribute handlerAttribute
	KeyMap          KeyMap
}

type HandlerRow struct {
	attributes map[handlerAttribute]*textinput.Model
}

func NewModel() *Model {
	m := &Model{
		ChosenHandler:   0,
		ChosenAttribute: MethodAttr,
		KeyMap:          DefaultKeyMap,
	}
	m.AppendHandlerRow(true)
	return m
}

type handlerAttribute int

const (
	MethodAttr handlerAttribute = iota
	PathAttr   handlerAttribute = iota
	EndAttr    handlerAttribute = iota
)

var handlerAttributeNames = map[handlerAttribute]string{
	MethodAttr: "Method",
	PathAttr:   "Path",
}

var attrSuggestions = map[handlerAttribute][]string{
	MethodAttr: {"GET", "POST", "PUT", "PATCH", "DELETE"},
	PathAttr: {
		"/api/v1/",
		"/internal/",
		"/ping",
		"/health",
		"api/v1/",
		"internal/",
		"ping",
		"health",
	},
}

func (m *Model) AppendHandlerRow(needFocus bool) {
	h := &HandlerRow{attributes: make(map[handlerAttribute]*textinput.Model)}
	for attrKey, attrName := range handlerAttributeNames {
		attr := textinput.New()
		attr.Placeholder = attrName
		attr.Prompt = ""
		attr.PromptStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
		attr.Cursor.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
		if needFocus {
			attr.Focus()
			needFocus = false
		}
		switch attrKey {
		case MethodAttr:
			attr.CharLimit = 20
			attr.Width = 20
		case PathAttr:
			attr.CharLimit = 45
			attr.Width = 45
		}
		if suggestions, ok := attrSuggestions[attrKey]; ok {
			attr.ShowSuggestions = true
			attr.SetSuggestions(suggestions)
		}
		attr.KeyMap.CharacterForward = m.KeyMap.CharacterForward
		attr.KeyMap.CharacterBackward = m.KeyMap.CharacterBackward

		h.attributes[attrKey] = &attr
	}
	m.Handlers = append(m.Handlers, h)
}

func (m *Model) Keys() help.KeyMap {
	return m.KeyMap
}

func (m *Model) ChooseHandler(newChosenHandler int) tea.Cmd {
	return m.Choose(newChosenHandler, m.ChosenAttribute)
}

func (m *Model) ChooseAttr(newChosenAttribute handlerAttribute) tea.Cmd {
	return m.Choose(m.ChosenHandler, newChosenAttribute)
}

func (m *Model) Choose(newChosenHandler int, newChosenAttribute handlerAttribute) tea.Cmd {
	if m.ChosenHandler != -1 {
		m.Handlers[m.ChosenHandler].attributes[m.ChosenAttribute].Blur()
	}
	m.ChosenHandler = newChosenHandler
	m.ChosenAttribute = newChosenAttribute
	if m.ChosenHandler != -1 {
		return m.Handlers[m.ChosenHandler].attributes[m.ChosenAttribute].Focus()
	}
	return nil
}

func (m *Model) moveToNextField() tea.Cmd {
	if m.ChosenAttribute < EndAttr-1 {
		return m.ChooseAttr(m.ChosenAttribute + 1)
	} else {
		if m.ChosenHandler == len(m.Handlers)-1 {
			m.AppendHandlerRow(false)
		}
		return m.Choose(
			m.ChosenHandler+1,
			0,
		)
	}
}

func (m *Model) IsEditable() bool {
	return m.ChosenHandler != -1
}

func (m *Model) SetNotEditable() {
	m.Choose(-1, 0)
}

func (m *Model) SetEditable() {
	m.Choose(0, 0)
}

func (m *Model) Update(msg tea.Msg) tea.Cmd {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.KeyMap.Up):
			return m.ChooseHandler(max(-1, m.ChosenHandler-1))
		case key.Matches(msg, m.KeyMap.Down):
			if m.ChosenHandler == len(m.Handlers)-1 {
				m.AppendHandlerRow(false)
			}
			return m.ChooseHandler(m.ChosenHandler + 1)
		case key.Matches(msg, m.KeyMap.PrevField):
			if m.ChosenAttribute > 0 {
				return m.ChooseAttr(m.ChosenAttribute - 1)
			} else if m.ChosenHandler > 0 {
				return m.Choose(m.ChosenHandler-1, EndAttr-1)
			}
			return nil
		case key.Matches(msg, m.KeyMap.NextField):
			return m.moveToNextField()
		case key.Matches(msg, m.KeyMap.AutoComplete):
			cmds := make([]tea.Cmd, 0)
			curInput := m.Handlers[m.ChosenHandler].attributes[m.ChosenAttribute]
			input, cmd := curInput.Update(msg)
			cmds = append(cmds, cmd)
			m.Handlers[m.ChosenHandler].attributes[m.ChosenAttribute] = &input
			if m.ChosenAttribute == MethodAttr {
				cmds = append(cmds, m.moveToNextField())
			}
			return tea.Batch(cmds...)
		}
	}

	if m.ChosenHandler != -1 {
		curInput := m.Handlers[m.ChosenHandler].attributes[m.ChosenAttribute]
		input, cmd := curInput.Update(msg)
		m.Handlers[m.ChosenHandler].attributes[m.ChosenAttribute] = &input
		return cmd
	}
	return nil
}

func (m Model) IsOk() bool {
	return true
}

func (m Model) Name() string {
	return "Service handlers"
}
