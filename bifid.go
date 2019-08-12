package caesar

// Bifid represents the Bifid cipher
// and conforms to the Cipher interface
// https://en.wikipedia.org/wiki/Bifid_cipher
type Bifid struct{}

// Encrypt operates on a plaintext string and a key string
// The function constructs an alphabet square from the key,
// and obtains substitution values from it. The substitution values
// are transposed and the transposed values are converted back into letters
func (a Bifid) Encrypt(plaintext, key string) (string, error) {
	return "", nil
}

// Decrypt operates on a ciphertext string and a key string
// The function constructs and alphabet square from the key,
// and obtains substitution values from it. The substitution values
// are de-transposed into values that are looked up in the square
// to obtain the original plaintext string.
func (a Bifid) Decrypt(ciphertext, key string) (string, error) {
	return "", nil
}
