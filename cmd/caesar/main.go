package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/stripedpajamas/caesar"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "caesar"
	app.Usage = "encrypt/decrypt using Caesar cipher"

	app.Commands = []cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "encrypt text",
			Action:  encrypt,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Value: "m",
					Usage: "the Caesar key",
				},
				cli.StringFlag{
					Name:  "text, t",
					Usage: "the text to encrypt",
				},
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "decrypt text",
			Action:  decrypt,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key, k",
					Value: "m",
					Usage: "the Caesar key",
				},
				cli.StringFlag{
					Name:  "text, t",
					Usage: "the text to encrypt",
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func encrypt(c *cli.Context) error {
	fmt.Println(caesar.Encrypt(c.String("text"), parseKey(c.String("key"))))
	return nil
}

func decrypt(c *cli.Context) error {
	fmt.Println(caesar.Decrypt(c.String("text"), parseKey(c.String("key"))))
	return nil
}

func parseKey(input string) int {
	if len(input) < 1 {
		log.Fatal("key must be a letter or a number")
	}
	k, err := strconv.Atoi(input)
	if err == nil {
		// input is a number
		return k
	}

	// input is a letter
	return int(strings.ToLower(input)[0] - 97)
}
