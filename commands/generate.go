package commands

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

func GenerateCommand() *cli.Command {
	return &cli.Command{
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
			err := generateAction(c)
			return err
		},
	}
}

func generateAction(c *cli.Context) error {
	size := c.Int("bytes")
	copy := c.Bool("copy")
	if size < 0 {
		fmt.Println("Invalid amount of bytes, defaulting to 32.")
		size = 32
	}

	key := make([]byte, size)
	_, err := rand.Read(key)
	if err != nil {
		return err
	}
	str := fmt.Sprintf("%x", key)
	err = clipboard.Init()
	if err != nil {
		log.Fatal(err)
	}
	if copy {
		clipboard.Write(clipboard.FmtText, []byte(str))
	}

	fmt.Println(str)
	return nil
}
