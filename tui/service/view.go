package service

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/dvordrova/gop/tui"
)

var (
	inactiveTabBorder = tabBorderWithBottom("┴", "─", "┴")
	activeTabBorder   = tabBorderWithBottom("┘", " ", "└")
	docStyle          = lipgloss.NewStyle().Padding(1, 2, 1, 2)
	highlightColor    = lipgloss.AdaptiveColor{Light: "#874BFD", Dark: "#7D56F4"}
	inactiveTabStyle  = lipgloss.NewStyle().Border(inactiveTabBorder, true).BorderForeground(highlightColor).Padding(0, 1)
	activeTabStyle    = inactiveTabStyle.Copy().Border(activeTabBorder, true)
	windowStyle       = lipgloss.NewStyle().BorderForeground(highlightColor).Padding(2, 0).Align(lipgloss.Left).Border(lipgloss.NormalBorder()).UnsetBorderTop()
)

func (m Model) View() string {
	view := strings.Builder{}
	view.WriteString("Model\n")
	var renderedTabs []string
	curLayout := m.Layouts[m.currentLayoutInd]
	for i, layout := range m.Layouts {
		tabName := layout.Name()
		var style lipgloss.Style
		isFirst, isLast, isActive := i == 0, i == len(m.Layouts)-1, i == m.currentLayoutInd
		if isActive {
			style = activeTabStyle.Copy()
		} else {
			style = inactiveTabStyle.Copy()
		}
		border, _, _, _, _ := style.GetBorder()
		if isFirst && isActive {
			border.BottomLeft = "│"
		} else if isFirst && !isActive {
			border.BottomLeft = "├"
		} else if isLast && isActive {
			border.BottomRight = "│"
		} else if isLast && !isActive {
			border.BottomRight = "┤"
		}
		style = style.Border(border)
		if isActive && !curLayout.IsEditable() {
			tabName = "*" + tabName
			style.Background(highlightColor)
		}
		renderedTabs = append(renderedTabs, style.Render(tabName))
	}

	row := lipgloss.JoinHorizontal(lipgloss.Top, renderedTabs...)
	view.WriteString(row)
	view.WriteString("\n")
	view.WriteString(windowStyle.Width((lipgloss.Width(row) - windowStyle.GetHorizontalFrameSize())).Render(
		curLayout.View(),
	))
	view.WriteString("\n")
	view.WriteString(tui.HelpView(m.help, curLayout.Keys(), m.KeyMap))
	return view.String()
}
