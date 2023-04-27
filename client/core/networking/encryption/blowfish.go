package encryption

import (
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"

	"golang.org/x/crypto/blowfish"
)

func EncryptBlowfish(key []byte, plaintext []byte) ([]byte, error) {
	// Create a new Blowfish cipher with the given key
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Generate a random IV
	iv := make([]byte, blowfish.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Encrypt the message
	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// Prepend the IV to the ciphertext
	ciphertext = append(iv, ciphertext...)

	return ciphertext, nil
}

func DecryptBlowfish(key []byte, ciphertext []byte) ([]byte, error) {
	// Create a new Blowfish cipher with the given key
	block, err := blowfish.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Extract the IV from the ciphertext
	if len(ciphertext) < blowfish.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := ciphertext[:blowfish.BlockSize]
	ciphertext = ciphertext[blowfish.BlockSize:]

	// Decrypt the message
	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
