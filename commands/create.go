package commands

import (
	"fmt"
	"io/ioutil"
	"xdarkyne/calyx/templates"

	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
)

func CreateCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "generate files automagically",
		Aliases: []string{
			"c",
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "file",
				Aliases: []string{
					"f",
				},
				Usage: "file name",
			},
			&cli.StringFlag{
				Name: "path",
				Aliases: []string{
					"p",
				},
				Usage: "destination path",
			},
		},
		Action: func(c *cli.Context) error {
			createAction(c)
			return nil
		},
	}
}

func createAction(c *cli.Context) {
	script := c.String("file")
	path := c.String("path")

	str, err := templates.Tmpl.ReadFile(fmt.Sprintf("%s.yaml", script))
	if err != nil {
		fmt.Println(err)
		return
	}
	var data templates.Template
	err = yaml.Unmarshal([]byte(str), &data)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch data.Type {
	case "create":
		str, err := templates.Tmpl.ReadFile(fmt.Sprintf("files/%s", data.File.SourceName))
		if err != nil {
			fmt.Println(err)
			return
		}
		if path == "" {
			path = fmt.Sprintf("./%s", data.File.DestinationName)
		}
		err = ioutil.WriteFile(path, str, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
