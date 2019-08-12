package caesar

import (
	"errors"
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

var adfgxReverse = map[byte]int{
	'A': 0,
	'D': 1,
	'F': 2,
	'G': 3,
	'X': 4,
}
var adfgx = [5]string{"A", "D", "F", "G", "X"}

// ADFGX represents the ADFGX cipher
// and conforms to the Cipher interface
// https://en.wikipedia.org/wiki/ADFGVX_cipher
type ADFGX struct{}

// Encrypt operates on a plaintext string and a key string
// that consists of two keys delimited by a semicolon (,).
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

	tb := newTranspositionBlock(k2)
	fractionated := tb.transpose(substitution.String())
	return fractionated, nil
}

// Decrypt operates on a ciphertext string and a key string
// that consists of two keys delimited by a semicolon (,).
// The function first transposes the letters according to key2,
// and then undoes the substitution using an alphabet square and key1.
func (a ADFGX) Decrypt(ciphertext, key string) (string, error) {
	input := runes.Clean(ciphertext)
	k1, k2, err := a.parseKeys(key)
	if err != nil {
		return "", err
	}

	tb := newTranspositionBlock(k2)
	unfranctionated := tb.detranspose(input)

	if len(unfranctionated)%2 != 0 {
		return "", errors.New("invalid ciphertext length")
	}

	// convert pairs of letters into coordinates in keyblock
	kb := newKeyblock(k1)
	var out strings.Builder
	for i := 0; i < len(unfranctionated); i += 2 {
		row, found := adfgxReverse[unfranctionated[i]]
		col, found := adfgxReverse[unfranctionated[i+1]]
		if !found {
			return "", errors.New("invalid ciphertext")
		}
		r, err := kb.getValue(location{row, col})
		if err != nil {
			return "", errors.New("invalid ciphertext")
		}
		out.WriteRune(r)
	}

	return out.String(), nil
}

func (a ADFGX) parseKeys(input string) (string, string, error) {
	split := strings.Split(input, ",")
	if len(split) < 2 {
		return "", "", errors.New("could not find two keys delimited by , in key input")
	}
	k1, k2 := runes.Clean(split[0]), runes.Clean(split[1])
	if len(k1) < 1 || len(k2) < 1 {
		return "", "", errors.New("found empty keys")
	}
	return k1, k2, nil
}
