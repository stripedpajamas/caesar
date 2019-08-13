package caesar

import (
	"errors"
	"strings"

	"github.com/stripedpajamas/caesar/runes"
)

// Bifid represents the Bifid cipher
// and conforms to the Cipher interface
// https://en.wikipedia.org/wiki/Bifid_cipher
type Bifid struct{}

// Encrypt operates on a plaintext string and a key string
// The function constructs an alphabet square from the key,
// and obtains substitution values from it. The substitution values
// are transposed and the transposed values are converted back into letters
func (a Bifid) Encrypt(plaintext, key string) (string, error) {
	kb := newKeyblock(key)
	cleanPlaintext := runes.Clean(plaintext)
	staging := make([]location, len(cleanPlaintext))

	// map plaintext runes to pairs in square
	for i, r := range cleanPlaintext {
		if !runes.IsLetter(r) {
			continue
		}
		loc, err := kb.getLocation(runes.ToUpper(r))
		if err != nil {
			// somehow a letter that isn't in the keyblock
			// skip it
			continue
		}
		staging[i] = loc
	}

	// split apart substitution pairs
	transposed := make([]int, len(staging)*2)

	for i := 0; i < len(staging); i++ {
		transposed[i] = staging[i].row
		transposed[i+len(staging)] = staging[i].col
	}

	// read off pairs of transposed int stream and
	// convert back to letters in square
	var out strings.Builder
	for i := 0; i < len(transposed); i += 2 {
		loc := location{row: transposed[i], col: transposed[i+1]}
		r, err := kb.getValue(loc)
		if err != nil {
			// this is an illegal state
			return "", errors.New("error converting transposed pairs into letters")
		}
		out.WriteRune(r)
	}

	return out.String(), nil
}

// Decrypt operates on a ciphertext string and a key string
// The function constructs and alphabet square from the key,
// and obtains substitution values from it. The substitution values
// are de-transposed into values that are looked up in the square
// to obtain the original plaintext string.
func (a Bifid) Decrypt(ciphertext, key string) (string, error) {
	kb := newKeyblock(key)
	cleanCiphertext := runes.Clean(ciphertext)
	staging := make([]location, len(cleanCiphertext))

	// map ciphertext runes to pairs in square
	for i, r := range cleanCiphertext {
		if !runes.IsLetter(r) {
			continue
		}
		loc, err := kb.getLocation(runes.ToUpper(r))
		if err != nil {
			// somehow a letter that isn't in the keyblock
			// skip it
			continue
		}
		staging[i] = loc
	}

	// read new pairs from staging (i, i + len)
	pairs := make([]location, len(staging))
	half := len(staging) / 2
	pairIdx := 0
	for i := 0; i < half; i++ {
		pairs[pairIdx] = location{row: staging[i].row, col: staging[i+half].row}
		pairs[pairIdx+1] = location{row: staging[i].col, col: staging[i+half].col}
		pairIdx += 2
	}

	// convert pairs into letters
	var out strings.Builder
	for _, loc := range pairs {
		r, err := kb.getValue(loc)
		if err != nil {
			// this is an illegal state
			return "", errors.New("error converting transposed pairs into letters")
		}
		out.WriteRune(r)
	}

	return out.String(), nil
}