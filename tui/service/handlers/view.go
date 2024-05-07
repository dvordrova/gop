package handlers

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var attrStyles = map[handlerAttribute]lipgloss.Style{
	MethodAttr: lipgloss.NewStyle().Width(21).Foreground(lipgloss.Color("63")),
	PathAttr:   lipgloss.NewStyle().Width(46).Foreground(lipgloss.Color("63")),
}

func (m Model) View() string {
	view := strings.Builder{}
	for _, handler := range m.Handlers {
		for attr := range EndAttr {
			style := attrStyles[attr]
			view.WriteString(style.Render(handler.attributes[attr].View()))
		}
		view.WriteString("\n")
	}
	return view.String()
}
