package caesar

// Vigenere represents the polyalphabetic substitution
// Vigenere cipher and conforms to the Cipher interface
// https://en.wikipedia.org/wiki/Vigen%C3%A8re_cipher
type Vigenere struct{}

func (v Vigenere) Encrypt(plaintext, key string) (string, error) {
	return "", nil
}

func (v Vigenere) Decrypt(ciphertext, key string) (string, error) {
	return "", nil
}
