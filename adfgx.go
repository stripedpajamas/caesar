package caesar

import (
	"errors"
	"fmt"
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

var adfgx = [5]string{"A", "D", "F", "G", "X"}

// ADFGX represents the ADFGX cipher
// and conforms to the Cipher interface
// https://en.wikipedia.org/wiki/ADFGVX_cipher
type ADFGX struct{}

// Encrypt operates on a plaintext string and a key string
// that consists of two keys delimited by a semicolon (;).
// The function constructs an alphabet square from key 1,
// and obtains substitution values from it. Then key 2 is used
// to transpose the values into the finished ciphertext.
func (a ADFGX) Encrypt(plaintext, key string) (string, error) {
	k1, k2, err := a.parseKeys(key)
	if err != nil {
		return "", err
	}
	kb := newKeyblock(k1)

	var substitution strings.Builder
	for _, r := range plaintext {
		if !runes.IsLetter(r) {
			continue
		}
		r, c, err := kb.getCoordinates(runes.ToUpper(r))
		if err != nil {
			// somehow a letter that isn't in the keyblock
			// skip it
			continue
		}
		substitution.WriteString(adfgx[r] + adfgx[c])
	}

	fmt.Println(newTranspositionBlock(substitution.String(), k2))
	return substitution.String(), nil
}

// Decrypt operates on a ciphertext string and a key string
// that consists of two keys delimited by a semicolon (;).
// The function first transposes the letters according to key2,
// and then undoes the substitution using an alphabet square and key1.
func (a ADFGX) Decrypt(plaintext, key string) (string, error) {
	return "", nil
}

func (a ADFGX) parseKeys(input string) (string, string, error) {
	split := strings.Split(input, ";")
	if len(split) < 2 {
		return "", "", errors.New("could not find two keys delimited by ; in key input")
	}

	var key1, key2 strings.Builder

	for _, r := range split[0] {
		if !runes.IsLetter(r) {
			continue
		}
		key1.WriteRune(r)
	}
	for _, r := range split[1] {
		if !runes.IsLetter(r) {
			continue
		}
		key2.WriteRune(r)
	}

	k1, k2 := key1.String(), key2.String()
	if len(k1) < 1 || len(k2) < 1 {
		return "", "", errors.New("found empty keys")
	}
	return k1, k2, nil
}
