package cipher

// Cipher represents an old school cipher
// with encryption and decryption methods
// that operate on ASCII inputs and keys
type Cipher interface {
	Encrypt(string, string) (string, error)
	Decrypt(string, string) (string, error)
}
