package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	ssmsession "github.com/chrispruitt/ssm-session-cli/lib"
	"github.com/spf13/cobra"
)

var (
	ssmSessionInput ssmsession.SsmSessionInput
)

func init() {
	log.SetFlags(0)

	rootCmd.AddCommand(SsmSessionCmd)
}

var SsmSessionCmd = &cobra.Command{
	Use:   "start",
	Short: "Start an ssm session on an ec2 instance using a selectable prompt",
	Run: func(cmd *cobra.Command, args []string) {
		promptInstance()

		err := ssmsession.StartSsmSession(&ssmSessionInput)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func promptInstance() {
	instances, err := ssmsession.GetInstances()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	clusterPrompt := &survey.Select{
		Message: "Select an instance:",
		Options: instances,
	}

	err = survey.AskOne(clusterPrompt, &ssmSessionInput.InstanceTemplate, survey.WithPageSize(20))
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	ssmSessionInput.InstanceId = strings.Split(ssmSessionInput.InstanceTemplate, " ")[0]
}
