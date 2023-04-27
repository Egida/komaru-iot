package encryption

import (
	"encoding/base64"
)

func EncodeBase64(input []byte) []byte {
	return []byte(base64.StdEncoding.EncodeToString(input))
}

func DecodeBase64(input []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(input))
}
