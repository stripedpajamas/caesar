package caesar

import (
	"strings"
)

const upperMin = 65
const upperMax = 90
const lowerMin = 97
const lowerMax = 122

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
		if isLetter(r) {
			out.WriteRune(shift(r, shiftVal))
		} else {
			out.WriteRune(r)
		}
	}
	return out.String()
}

func shift(r rune, n int) rune {
	var top, bottom rune
	if isLower(r) {
		top = lowerMax
		bottom = lowerMin
	} else {
		top = upperMax
		bottom = upperMin
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

func isLetter(r rune) bool {
	return isLower(r) || isUpper(r)
}

func isUpper(r rune) bool {
	return (r >= upperMin && r <= upperMax)
}

func isLower(r rune) bool {
	return (r >= lowerMin && r <= lowerMax)
}
