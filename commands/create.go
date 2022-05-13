package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
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

	for _, action := range data.Actions {
		fmt.Println(action.Name)
		switch action.Type {
		case "create":
			str, err := templates.Tmpl.ReadFile(fmt.Sprintf("files/%s", action.File.SourceName))
			if err != nil {
				fmt.Println(err)
				return
			}
			path := c.String("path")
			if path == "" {
				path = fmt.Sprintf("./%s", action.File.DestinationName)
			}
			err = ioutil.WriteFile(path, str, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
		case "delete":
			path := c.String("path")
			if path == "" {
				path = fmt.Sprintf("./%s", action.File.SourceName)
			}
			if _, err := os.Stat(path); err == nil {
				// path/to/whatever exists
				err = os.Remove(path)
				if err != nil {
					fmt.Println(err)
					return
				}
			} else if errors.Is(err, os.ErrNotExist) {
				// path/to/whatever does *not* exist
				fmt.Println(err)
			}
		}
	}

}
