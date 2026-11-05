package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	header := titleStyle.Render("AWS Instance Manager")

	footer := footerStyle.Render(
		"↑/↓ Navigate  •  / Search  •  Enter Select  •  q Quit",
	)

	status := statusStyle.Render(m.status)
	v := tea.NewView(fmt.Sprintf(
		"%s\n%s\n\n%s\n%s",
		header,
		m.list.View(),
		status,
		footer,
	))
	v.AltScreen = true
	return v
}
