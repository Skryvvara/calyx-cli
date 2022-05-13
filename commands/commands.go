package commands

import (
	"github.com/urfave/cli/v2"
)

// Commands is a list of CLI commands
func InitializeCommands() []*cli.Command {
	return []*cli.Command{
		CreateCommand(),
		GenerateCommand(),
	}
}
