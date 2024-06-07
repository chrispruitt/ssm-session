package ecs

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type SsmSessionInput struct {
	InstanceId       string
	InstanceTemplate string
}

func GetInstances() ([]string, error) {

	instances := []*ec2.Instance{}
	instancesTemplates := []string{}

	pageNum := 0

	err := ec2Client.DescribeInstancesPages(&ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running")},
			},
		},
	},
		func(page *ec2.DescribeInstancesOutput, b bool) bool {
			pageNum++
			for _, res := range page.Reservations {
				instances = append(instances, res.Instances...)
				for _, i := range res.Instances {
					name := ""
					for _, tag := range i.Tags {
						if *tag.Key == "Name" {
							name = *tag.Value
						}
					}
					instancesTemplates = append(instancesTemplates, fmt.Sprintf("%s %s", *i.InstanceId, name))
				}
			}
			return pageNum <= 100
		})
	return instancesTemplates, err
}

// aws ssm start-session --target $INSTANCE_ID
func StartSsmSession(input *SsmSessionInput) error {
	args := []string{
		"ssm",
		"start-session",
		"--target",
		input.InstanceId,
	}

	if err := runCommand("aws", args...); err != nil {
		return err
	}
	return nil
}

func runCommand(process string, args ...string) error {
	cmd := exec.Command(process, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGINT)
	go func() {
		for {
			select {
			case <-sigs:
			}
		}
	}()
	defer close(sigs)

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
