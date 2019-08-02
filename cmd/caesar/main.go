package main

import (
	"errors"
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
			Usage: "the cipher to use (one of caesar, playfair, vigenere)",
		},
		cli.StringFlag{
			Name:  "key, k",
			Usage: "the key to encrypt/decrypt with",
		},
		cli.StringFlag{
			Name:  "text, t",
			Usage: "the text to encrypt",
		},
		cli.IntFlag{
			Name:  "groups, g",
			Value: 5,
			Usage: "how many letters in each printed output group",
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
					c.GlobalInt("groups"),
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
					0,
				)
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func handle(operation, cipherType, input, key string, groups int) error {
	cipher, err := getCipher(cipherType)

	if err != nil {
		return err
	}

	var output string
	switch operation {
	case "encrypt":
		output, err = encrypt(cipher, input, key)
	case "decrypt":
		output, err = decrypt(cipher, input, key)
	}

	if err != nil {
		return err
	}

	print(output, groups)
	return nil
}

func getCipher(cipher string) (caesar.Cipher, error) {
	switch strings.ToLower(cipher) {
	case "caesar":
		return caesar.Caesar{}, nil
	case "playfair":
		return caesar.Playfair{}, nil
	case "vigenere":
		return caesar.Vigenere{}, nil
	}
	return nil, errors.New("unrecognized cipher type")
}

func encrypt(cipher caesar.Cipher, plaintext, key string) (string, error) {
	return cipher.Encrypt(plaintext, key)
}

func decrypt(cipher caesar.Cipher, ciphertext, key string) (string, error) {
	return cipher.Decrypt(ciphertext, key)
}

func print(output string, groups int) {
	if groups == 0 {
		fmt.Println(output)
		return
	}
	var out strings.Builder
	for idx, r := range output {
		if idx != 0 && idx%groups == 0 {
			out.WriteRune(' ')
		}
		out.WriteRune(r)
	}
	fmt.Println(out.String())
}
