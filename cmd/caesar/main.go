package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stripedpajamas/caesar"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "caesar"
	app.Usage = "encrypt/decrypt using oldschool ciphers"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "cipher, c",
			Usage: "the cipher to use (one of caesar, playfair)",
		},
		cli.StringFlag{
			Name:  "key, k",
			Usage: "the key to encrypt/decrypt with",
		},
		cli.StringFlag{
			Name:  "text, t",
			Usage: "the text to encrypt",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "encrypt",
			Aliases: []string{"e"},
			Usage:   "encrypt text",
			Action: func(c *cli.Context) error {
				return handle(
					"encrypt",
					c.GlobalString("cipher"),
					c.GlobalString("text"),
					c.GlobalString("key"),
				)
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "decrypt text",
			Action: func(c *cli.Context) error {
				return handle(
					"decrypt",
					c.GlobalString("cipher"),
					c.GlobalString("text"),
					c.GlobalString("key"),
				)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func handle(operation, cipherType, input, key string) error {
	var c caesar.Cipher
	switch strings.ToLower(cipherType) {
	case "caesar":
		c = caesar.Caesar{}
	case "playfair":
		c = caesar.Playfair{}
	}

	var output string
	var err error
	switch operation {
	case "encrypt":
		output, err = encrypt(c, input, key)
	case "decrypt":
		output, err = decrypt(c, input, key)
	}

	if err != nil {
		return err
	}

	print(output)
	return nil
}

func encrypt(cipher caesar.Cipher, plaintext, key string) (string, error) {
	return cipher.Encrypt(plaintext, key)
}

func decrypt(cipher caesar.Cipher, ciphertext, key string) (string, error) {
	return cipher.Decrypt(ciphertext, key)
}

func print(output string) {
	// TODO accept print options and apply them here
	fmt.Println(output)
}
