package encryption

import (
	"fmt"
)

func Encrypt(blowfishKey, chachaKey, plainText []byte) ([]byte, error) {

	blowfish, err := EncryptBlowfish(blowfishKey, plainText)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 0), err
	}

	chacha, err := EncryptChaCha20(chachaKey, blowfish)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 0), err
	}

	return EncodeBase64(chacha), err
}

func Decrypt(blowfishKey, chachaKey, cipherText []byte) ([]byte, error) {

	newText, _ := DecodeBase64(cipherText)

	chacha, err := DecryptChaCha20(chachaKey, newText)
	if err != nil {
		return make([]byte, 0), err
	}

	blowfish, err := DecryptBlowfish(blowfishKey, chacha)
	if err != nil {
		return make([]byte, 0), err
	}

	return blowfish, err
}
