package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"github.com/glhou/aws-ec2/internal/aws"
	"github.com/glhou/aws-ec2/internal/tui"
)

func main() {

	regions := []string{"us-east-1", "ap-northeast-1"}

	instances, err := aws.ListInstances(regions)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	p := tea.NewProgram(
		tui.NewModel(instances),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}
