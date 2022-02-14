package actions

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/urfave/cli/v2"
	"golang.design/x/clipboard"
)

func Generate(c *cli.Context) error {
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
	if copy == true {
		clipboard.Write(clipboard.FmtText, []byte(str))
	}

	fmt.Println(str)
	return nil
}
