package aws

import "os/exec"

func SSMConnect(instance Instance) *exec.Cmd {
	return exec.Command(
		"aws",
		"ssm",
		"start-session",
		"--target", instance.ID,
		"--region", instance.Region,
	)
}
