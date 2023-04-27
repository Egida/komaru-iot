package encryption

import (
	chacha20 "golang.org/x/crypto/chacha20"
)

func EncryptChaCha20(key []byte, data []byte) ([]byte, error) {
	cipherBlock, err := chacha20.NewUnauthenticatedCipher(key, make([]byte, chacha20.NonceSize))
	if err != nil {
		return make([]byte, 0), err
	}

	ciphertext := make([]byte, len(data))
	cipherBlock.XORKeyStream(ciphertext, data)
	return ciphertext, nil
}

func DecryptChaCha20(key []byte, data []byte) ([]byte, error) {
	cipherBlock, err := chacha20.NewUnauthenticatedCipher(key, make([]byte, chacha20.NonceSize))
	if err != nil {
		return make([]byte, 0), err
	}
	decrypted := make([]byte, len(data))
	cipherBlock, err = chacha20.NewUnauthenticatedCipher(key, make([]byte, chacha20.NonceSize))
	if err != nil {
		panic(err)
	}
	cipherBlock.XORKeyStream(decrypted, data)
	return decrypted, nil
}
