package actions

import (
	"github.com/urfave/cli/v2"
)

// Commands is a list of CLI commands
func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "status",
			Usage: "show the status of the cli",
			Aliases: []string{
				"s",
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "status",
					Usage: "status of the app",
				},
			},
			Action: func(c *cli.Context) error {
				StatusCommand(c)
				return nil
			},
		},
		{
			Name:  "generate",
			Usage: "generate things",
			Aliases: []string{
				"gen",
				"g",
			},
			Flags: []cli.Flag{
				&cli.IntFlag{
					Name:  "bytes",
					Usage: "set the amount of bytes",
					Aliases: []string{
						"b",
					},
					Value: 32,
				},
				&cli.BoolFlag{
					Name:  "copy",
					Usage: "copy the output string to clipbard",
					Aliases: []string{
						"c",
					},
					Value: false,
				},
			},
			Action: func(c *cli.Context) error {
				err := Generate(c)
				return err
			},
		},
	}
}
