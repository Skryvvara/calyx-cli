package main

import (
	"log"
	"os"
	"xdarkyne/calyx/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "calyx",
		Usage: "say hello",
	}

	app.Commands = commands.InitializeCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
