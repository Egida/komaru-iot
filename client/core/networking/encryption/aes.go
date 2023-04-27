package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	realRand "math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func Random(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[realRand.Intn(len(letterRunes))]
	}
	return string(b)
}

func EncryptAes(key []byte, plaintext []byte) ([]byte, error) {
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return make([]byte, 0), err
	}

	block := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, block); err != nil {
		return make([]byte, 0), err
	}

	paddedPlainText := padAes(plaintext, aes.BlockSize)
	cipherText := make([]byte, len(paddedPlainText))

	cbcEncrypter := cipher.NewCBCEncrypter(aesCipher, block)
	cbcEncrypter.CryptBlocks(cipherText, paddedPlainText)

	return append(block, cipherText...), nil
}

func DecryptAes(key []byte, cipherText []byte) ([]byte, error) {
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	block := cipherText[:aes.BlockSize]
	cipherTextData := cipherText[aes.BlockSize:]
	plainText := make([]byte, len(cipherTextData))

	cbcDecrypter := cipher.NewCBCDecrypter(aesCipher, block)
	cbcDecrypter.CryptBlocks(plainText, cipherTextData)

	unpaddedPlainText, err := unpadAes(plainText)
	if err != nil {
		return nil, err
	}

	return unpaddedPlainText, nil
}

func GenerateAesKey() ([]byte, error) {
	realRand.Seed(time.Now().UnixNano())
	return []byte(Random(16)), nil
}

func padAes(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func unpadAes(data []byte) ([]byte, error) {
	length := len(data)
	unpadding := int(data[length-1])
	if length < int(data[length-1]) {
		return nil, errors.New("invalid padding")
	}
	return data[:length-unpadding], nil
}
