package main

import (
	"github.com/chrispruitt/ssm-session-cli/cmd"
)

// version of ssm-session-cli. Overwritten during build
var version = "development"

func main() {
	cmd.Execute(version)
}
