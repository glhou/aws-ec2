package aws

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Instance struct {
	Name   string `json:"Name"`
	ID     string `json:"ID"`
	Region string `json:"Region"`
	IP     string `json:"IP"`
	Status string `json:"Status"`
}

func (i Instance) Title() string {
	return fmt.Sprintf("%-20s %s", i.Name, i.Region)
}

func (i Instance) Description() string {
	return fmt.Sprintf("%s  •  %s  •  %s", i.ID, i.IP, i.Status)
}

func (i Instance) FilterValue() string {
	return i.Name + " " + i.ID + " " + i.Region + " " + i.IP
}

func ListInstances(regions []string) ([]Instance, error) {
	var totalInstances []Instance
	for _, region := range regions {
		out, err := exec.Command(
			"aws",
			"ec2",
			"describe-instances",
			"--region",
			region,
			"--output",
			"json",
			"--query",
			"Reservations[].Instances[].{Name: Tags[?Key==`Name`]|[0].Value, ID: InstanceId, IP: PrivateIpAddress, Status: State.Name}",
			"--no-cli-pager",
		).Output()
		if err != nil {
			return nil, err
		}
		var instances []Instance

		if err := json.Unmarshal(out, &instances); err != nil {
			fmt.Println("Couldn't unmarshal:", out)
			return nil, err
		}

		for _, instance := range instances {
			instance.Region = region
			totalInstances = append(totalInstances, instance)
		}
	}
	return totalInstances, nil
}
