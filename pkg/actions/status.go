package actions

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func StatusCommand() *cli.Command {
	return &cli.Command{
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
			statusAction(c)
			return nil
		},
	}
}

func statusAction(c *cli.Context) {
	status := c.String("status")
	fmt.Println(status)
}
