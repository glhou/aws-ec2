package tui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"github.com/glhou/aws-ec2/internal/aws"
)

type Model struct {
	list   list.Model
	status string
}

func NewModel(instances []aws.Instance) Model {
	items := make([]list.Item, len(instances))
	for i, instance := range instances {
		items[i] = instance
	}

	delegate := list.NewDefaultDelegate()
	l := list.New(items, delegate, 0, 0)

	l.Title = "EC2 Instances"
	l.SetShowStatusBar(false)
	l.SetShowHelp(false)
	l.SetFilteringEnabled(true)
	l.DisableQuitKeybindings()

	return Model{
		list:   l,
		status: "Select and instance to connect",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}
