package main

import (
	"log"
	"os"
	"xdarkyne/calyx/pkg/actions"

	"github.com/urfave/cli"
)

func main() {
	app := &cli.App{
		Name:  "calyx",
		Usage: "say hello",
	}

	app.Commands = actions.Commands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
