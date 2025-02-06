package lqcrypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"github.com/go-liquor/liquor-sdk/config"
	"golang.org/x/crypto/bcrypt"
)

type CryptoHelper struct {
	cost int
}

func NewCryptoHelper(cfg *config.Config) *CryptoHelper {
	return &CryptoHelper{
		cost: cfg.GetPassswordBcryptCost(),
	}
}

// Hash generates a bcrypt hash from the input data.
//
// Parameters:
//   - data: The string to hash
//
// Returns:
//   - string: The hashed string
//   - error: nil if successful, error otherwise
//
// Example:
//
//	hash, err := crypto.Hash("mypassword")
//	if err != nil {
//	}
func (c *CryptoHelper) Hash(data string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(data), c.cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

// CompareHash checks if the provided data matches the hash.
//
// Parameters:
//   - hash: The bcrypt hash to compare against
//   - data: The plain text to compare
//
// Returns:
//   - bool: true if the data matches the hash, false otherwise
//
// Example:
//
//	matches := crypto.CompareHash(storedHash, "mypassword")
func (c *CryptoHelper) CompareHash(hash, data string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil
}

// Encrypt encrypts data using AES-GCM.
//
// Parameters:
//   - data: The data to encrypt
//   - key: The encryption key (must be 16, 24, or 32 bytes for AES-128, AES-192, or AES-256)
//
// Returns:
//   - []byte: The encrypted data
//   - error: nil if successful, error otherwise
//
// Example:
//
//	key, _ := crypto.GenerateKey(256)
//	encrypted, err := crypto.Encrypt([]byte("secret data"), key)
func (c *CryptoHelper) Encrypt(data []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

// Decrypt decrypts AES-GCM encrypted data.
//
// Parameters:
//   - encrypted: The encrypted data
//   - key: The decryption key (must match the encryption key)
//
// Returns:
//   - []byte: The decrypted data
//   - error: nil if successful, error otherwise
//
// Example:
//
//	decrypted, err := crypto.Decrypt(encryptedData, key)
func (c *CryptoHelper) Decrypt(encrypted []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	if len(encrypted) < gcm.NonceSize() {
		return nil, errors.New("encrypted data too short")
	}

	nonce, ciphertext := encrypted[:gcm.NonceSize()], encrypted[gcm.NonceSize():]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}

// GenerateKey generates a random key of specified bits length.
//
// Parameters:
//   - bits: The key length in bits (typically 128, 192, or 256 for AES)
//
// Returns:
//   - []byte: The generated key
//   - error: nil if successful, error otherwise
//
// Example:
//
//	key, err := crypto.GenerateKey(256)
func (c *CryptoHelper) GenerateKey(bits int) ([]byte, error) {
	key := make([]byte, bits/8)
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		return nil, err
	}
	return key, nil
}
