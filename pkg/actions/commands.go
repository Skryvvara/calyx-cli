package actions

import (
	"github.com/urfave/cli"
)

// Commands is a list of CLI commands
func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:  "status",
			Usage: "show the status of the cli",
			Aliases: []string{
				"s",
			},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "status",
					Usage: "status of the app",
				},
			},
			Action: func(c *cli.Context) error {
				StatusCommand(c)
				return nil
			},
		},
	}
}
