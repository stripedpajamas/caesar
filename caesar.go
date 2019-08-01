package caesar

import (
	"errors"
	"strconv"
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

// Encrypt converts alphabetic characters to their down-shifted
// values based on the key parameter (e.g. a shifted 1 = b)
func Encrypt(plaintext string, key string) (string, error) {
	parsedKey, err := parseKey(key)
	if err != nil {
		return "", err
	}
	return process(plaintext, parsedKey%26), nil
}

// Decrypt converts alphabetic characters to their up-shifted
// values based on the key parameter (e.g. a shifted 1 = z)
func Decrypt(ciphertext string, key string) (string, error) {
	parsedKey, err := parseKey(key)
	if err != nil {
		return "", err
	}
	return process(ciphertext, -1*(parsedKey%26)), nil
}

func parseKey(input string) (int, error) {
	if len(input) < 1 {
		return 0, errors.New("key must be a letter or a number")
	}
	k, err := strconv.Atoi(input)
	if err == nil {
		// input is a number
		return k, nil
	}

	// input is a letter
	head := rune(input[0])
	if !runes.IsLetter(head) {
		return 0, errors.New("key must be a letter or a number")
	}
	return int(head - 97), nil
}

func process(input string, shiftVal int) string {
	var out strings.Builder
	for _, r := range input {
		if runes.IsLetter(r) {
			out.WriteRune(shift(r, shiftVal))
		} else {
			out.WriteRune(r)
		}
	}
	return out.String()
}

func shift(r rune, n int) rune {
	var top, bottom rune
	if runes.IsLower(r) {
		top = runes.LowerMax
		bottom = runes.LowerMin
	} else {
		top = runes.UpperMax
		bottom = runes.UpperMin
	}

	ret := r + rune(n)
	if ret > top {
		ret -= 26
	}
	if ret < bottom {
		ret += 26
	}

	return ret
}
