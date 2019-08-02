package caesar

import (
	"errors"
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

// Vigenere represents the polyalphabetic substitution
// Vigenere cipher and conforms to the Cipher interface
// https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher
type Vigenere struct{}

// Encrypt converts alphabetic characters to their corresponding
// values based on the key parameter
func (v Vigenere) Encrypt(plaintext, key string) (string, error) {
	err := v.validateKey(key)
	if err != nil {
		return "", err
	}

	return v.process(plaintext, key, false), nil
}

// Decrypt converts alphabetic characters to their corresponding
// values based on the key parameter
func (v Vigenere) Decrypt(ciphertext, key string) (string, error) {
	err := v.validateKey(key)
	if err != nil {
		return "", err
	}

	return v.process(ciphertext, key, true), nil
}

func (v Vigenere) validateKey(key string) error {
	for _, r := range key {
		if !runes.IsLetter(r) {
			return errors.New("key must be alphabetic")
		}
	}
	return nil
}

func (v Vigenere) process(input, key string, reverse bool) string {
	var diff int
	if reverse {
		diff = -1
	} else {
		diff = 1
	}
	// caesar cipher encrypt each letter of input
	// using the corresponding letter of the key as the caesar key
	var out strings.Builder
	keyIdx := 0
	for _, r := range input {
		if !runes.IsLetter(r) {
			// skip non-letters
			continue
		}
		k := runes.ToLower(rune(key[keyIdx%len(key)]))
		shiftVal := int(k - 97)
		out.WriteRune(runes.ToUpper(Caesar{}.shift(r, diff*shiftVal)))
		keyIdx++
	}
	return out.String()
}
