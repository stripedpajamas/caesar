package caesar

import (
	"errors"
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

const xPad = 'X'
const qPad = 'Q'
const alphabet = "ABCDEFGHIKLMNOPQRSTUVWXYZ"

// Playfair represents the Playfair cipher
// and conforms to the Cipher interface
// https://en.wikipedia.org/wiki/Playfair_cipher
type Playfair struct{}
type keyblock struct {
	block  [5][5]rune
	lookup map[rune]location
}

type location struct {
	row int
	col int
}

// Encrypt constructs a playfair alphabet from key,
// and then encrypts plaintext using it
func (pf Playfair) Encrypt(plaintext string, key string) (string, error) {
	kb := pf.newKeyblock(key)
	pt := pf.unrollPlaintext(plaintext)
	return pf.process(pt, kb, false), nil
}

// Decrypt constructs a playfair alphabet from key,
// and then decrypts ciphertext using it
func (pf Playfair) Decrypt(ciphertext string, key string) (string, error) {
	kb := pf.newKeyblock(key)
	if len(ciphertext)%2 != 0 {
		return "", errors.New("invalid ciphertext; length not even")
	}
	return pf.process(ciphertext, kb, true), nil
}

func (pf Playfair) process(in string, kb *keyblock, reverse bool) string {
	var out strings.Builder
	for i := 0; i < len(in); i += 2 {
		a, b := rune(in[i]), rune(in[i+1])
		x, y := kb.getCorrespondingPair(a, b, reverse)
		out.WriteRune(x)
		out.WriteRune(y)
	}

	return out.String()
}

// remove non-alphabetic chars;
// break message into digrams;
// pad repeat digrams;
// pad incomplete digrams
func (pf Playfair) unrollPlaintext(plaintext string) string {
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

func (pf Playfair) newKeyblock(key string) *keyblock {
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

// - If the letters appear on the same row of your table,
//   replace them with the letters to their immediate right respectively
// - If the letters appear on the same column of your table,
//   replace them with the letters immediately below respectively
// - If the letters are not on the same row or column,
//   replace them with the letters on the same row respectively
//   but at the other pair of corners of the rectangle defined by the original pair.
//   The order is important – the first letter of the encrypted pair is the
//   one that lies on the same row as the first letter of the plaintext pair.
func (kb *keyblock) getCorrespondingPair(a, b rune, reverse bool) (x, y rune) {
	var diff int
	if reverse {
		diff = -1
	} else {
		diff = 1
	}
	aLoc, bLoc := kb.lookup[a], kb.lookup[b]
	var xLoc, yLoc location
	if aLoc.row == bLoc.row {
		xLoc = location{row: aLoc.row, col: (aLoc.col + diff + 5) % 5}
		yLoc = location{row: bLoc.row, col: (bLoc.col + diff + 5) % 5}
	} else if aLoc.col == bLoc.col {
		xLoc = location{row: (aLoc.row + diff + 5) % 5, col: aLoc.col}
		yLoc = location{row: (bLoc.row + diff + 5) % 5, col: bLoc.col}
	} else {
		xLoc = location{row: aLoc.row, col: bLoc.col}
		yLoc = location{row: bLoc.row, col: aLoc.col}
	}
	x = kb.block[xLoc.row][xLoc.col]
	y = kb.block[yLoc.row][yLoc.col]
	return
}
