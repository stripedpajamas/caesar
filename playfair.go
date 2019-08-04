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

// Encrypt constructs a playfair alphabet from key,
// and then encrypts plaintext using it
func (pf Playfair) Encrypt(plaintext string, key string) (string, error) {
	kb := newKeyblock(key)
	pt := pf.unrollPlaintext(plaintext)
	return pf.process(pt, kb, false), nil
}

// Decrypt constructs a playfair alphabet from key,
// and then decrypts ciphertext using it
func (pf Playfair) Decrypt(ciphertext string, key string) (string, error) {
	if err := validateCiphertext(ciphertext); err != nil {
		return "", err
	}
	kb := newKeyblock(key)
	return pf.process(ciphertext, kb, true), nil
}

func validateCiphertext(ciphertext string) error {
	sum := 0
	for _, r := range ciphertext {
		if runes.IsLetter(r) {
			sum++
		}
	}
	if sum%2 != 0 {
		return errors.New("invalid ciphertext; length not even")
	}
	return nil
}

func (pf Playfair) process(input string, kb *keyblock, reverse bool) string {
	var tmp strings.Builder
	for _, r := range input {
		if runes.IsLetter(r) {
			tmp.WriteRune(r)
		}
	}
	in := tmp.String()
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
