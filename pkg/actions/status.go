package actions

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func StatusCommand(c *cli.Context) {
	status := c.String("status")
	fmt.Println(status)
}
