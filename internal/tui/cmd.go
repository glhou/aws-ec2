package tui

import (
	tea "charm.land/bubbletea/v2"
	"github.com/glhou/aws-ec2/internal/aws"
)

type sessionFinishedMsg struct {
	err error
}

func connect(instance aws.Instance) tea.Cmd {
	cmd := aws.SSMConnect(instance)

	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return sessionFinishedMsg{err: err}
	})
}
