package caesar

import (
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

// Encrypt converts alphabetic characters to their down-shifted
// values based on the key parameter (e.g. a shifted 1 = b)
func Encrypt(plaintext string, key int) string {
	return process(plaintext, key%26)
}

// Decrypt converts alphabetic characters to their up-shifted
// values based on the key parameter (e.g. a shifted 1 = z)
func Decrypt(ciphertext string, key int) string {
	return process(ciphertext, -1*(key%26))
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
