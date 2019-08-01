package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/stripedpajamas/caesar"
	"github.com/stripedpajamas/caesar/cipher"
	"github.com/stripedpajamas/caesar/playfair"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "caesar"
	app.Usage = "encrypt/decrypt using oldschool ciphers"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "cipher, c",
			Value: "the cipher to use (one of caesar, playfair)",
		},
		cli.StringFlag{
			Name:  "key, k",
			Value: "m",
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
					strings.ToLower(c.String("cipher")),
					c.String("text"),
					c.String("key"),
				)
			},
		},
		{
			Name:    "decrypt",
			Aliases: []string{"d"},
			Usage:   "decrypt text",
			Action: func(c *cli.Context) error {
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func handle(operation, cipherType, input, key string) error {
	var c cipher.Cipher
	switch strings.ToLower(cipherType) {
	case "caesar":
		c = caesar.Caesar{}
	case "playfair":
		c = playfair.Playfair{}
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

func encrypt(cipher cipher.Cipher, plaintext, key string) (string, error) {
	return cipher.Encrypt(plaintext, key)
}

func decrypt(cipher cipher.Cipher, ciphertext, key string) (string, error) {
	return cipher.Decrypt(ciphertext, key)
}

func print(output string) {
	// TODO accept print options and apply them here
	fmt.Println(output)
}
