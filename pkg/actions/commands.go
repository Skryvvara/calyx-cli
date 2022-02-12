package actions

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

// Commands is a list of CLI commands
func Commands() {
	app := &cli.App{
		Name:  "calyx",
		Usage: "say hello",
	}

	app.Commands = []cli.Command{
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

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
