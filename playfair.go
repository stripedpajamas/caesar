package caesar

import (
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

const xPad = 'X'
const qPad = 'Q'

// PlayfairEncrypt constructs a playfair alphabet from key,
// and then encrypts plaintext using it
func PlayfairEncrypt(plaintext string, key string) string {
	return unrollPlaintext(plaintext)
}

// PlayfairDecrypt constructs a playfair alphabet from key,
// and then decrypts ciphertext using it
func PlayfairDecrypt(ciphertext string, key string) string {
	return ""
}

// remove non-alphabetic chars;
// break message into digrams;
// pad repeat digrams;
// pad incomplete digrams
func unrollPlaintext(plaintext string) string {
	// for any letter in pt, if the next letter is same,
	// add an x between current and next and continue processing
	var out strings.Builder
	a, b := 0, 1
	for a < len(plaintext) && b <= len(plaintext) {
		aChar := rune(plaintext[a])
		var bChar rune
		if b == len(plaintext) {
			// odd length plaintext, pad to even
			bChar = xPad
		} else {
			bChar = rune(plaintext[b])
		}
		if !runes.IsLetter(aChar) {
			a++
			b++
			continue
		}
		if !runes.IsLetter(bChar) {
			b++
			continue
		}
		aChar = runes.ToUpper(aChar)
		bChar = runes.ToUpper(bChar)
		out.WriteRune(aChar)
		if aChar == bChar {
			// duplicate letter, add pad
			if aChar == xPad {
				out.WriteRune(qPad)
			} else {
				out.WriteRune(xPad)
			}
			a, b = b, b+1
			continue
		}
		out.WriteRune(bChar)
		a, b = b+1, b+2
	}

	return out.String()
}
