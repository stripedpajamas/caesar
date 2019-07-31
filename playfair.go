package caesar

import (
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

const xPad = 'X'
const qPad = 'Q'
const alphabet = "ABCDEFGHIKLMNOPQRSTUVWXYZ"

type keyblock struct {
	block  [5][5]rune
	lookup map[rune]location
}

type location struct {
	row int
	col int
}

func (kb *keyblock) getCorrespondingPair(a, b rune) {
	// TODO
}

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

func newKeyblock(key string) *keyblock {
	kb := new(keyblock)
	kb.block = [5][5]rune{}
	kb.lookup = make(map[rune]location)

	// get unique letters from key
	seen := make(map[rune]bool)
	uniqueKey, keyIdx := [25]rune{}, 0
	for _, r := range key {
		if !runes.IsLetter(r) || r == 'j' || r == 'J' {
			continue
		}
		ur := runes.ToUpper(r)
		if seen[ur] {
			continue
		}
		uniqueKey[keyIdx] = ur
		seen[ur] = true
		keyIdx++
	}

	// add the rest of the alphabet (excepting already used letters)
	for _, r := range alphabet {
		if seen[r] {
			continue
		}
		uniqueKey[keyIdx] = r
		keyIdx++
	}

	keyIdx = 0
	// write the resulting key into the keyblock
	for row := 0; row < 5; row++ {
		kb.block[row] = [5]rune{}
		for col := 0; col < 5; col++ {
			kb.block[row][col] = uniqueKey[keyIdx]
			kb.lookup[uniqueKey[keyIdx]] = location{row, col}
			keyIdx++
		}
	}

	return kb
}
