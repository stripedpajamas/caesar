package caesar

import (
	"errors"
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

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
	return "", nil
}

// Decrypt operates on a ciphertext string and a key string
// that consists of two keys delimited by a semicolon (;).
// The function constructs an alphabet square from key 1,
// and obtains substitution values from it. Then key 2 is used
// to transpose the values into the finished plaintext.
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
