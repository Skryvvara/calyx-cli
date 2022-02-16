package actions

import (
	"github.com/urfave/cli/v2"
)

// Commands is a list of CLI commands
func Commands() []*cli.Command {
	return []*cli.Command{
		StatusCommand(),
		GenerateCommand(),
	}
}
