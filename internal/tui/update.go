package tui

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
	"github.com/glhou/aws-ec2/internal/aws"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-5)

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "enter":
			selected, ok := m.list.SelectedItem().(aws.Instance)
			if !ok {
				return m, nil
			}
			return m, connect(selected)
		}
	case sessionFinishedMsg:
		if msg.err != nil {
			m.status = fmt.Sprintf("Session failed: %v", msg.err)
		} else {
			m.status = "Session ended"
		}
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
