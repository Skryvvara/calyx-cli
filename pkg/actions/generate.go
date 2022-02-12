package actions

import (
	"crypto/rand"
	"fmt"

	"github.com/urfave/cli/v2"
)

func Generate(c *cli.Context) error {
	size := c.Int("bytes")
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
	fmt.Println(str)
	return nil
}
