package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

func StatusCommand(c *cli.Context) {
	status := c.String("status")
	fmt.Println(status)
}
