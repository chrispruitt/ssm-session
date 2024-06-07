package main

import (
	"github.com/chrispruitt/ssm-session/cmd"
)

// version of ssm-session. Overwritten during build
var version = "development"

func main() {
	cmd.Execute(version)
}
