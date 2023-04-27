package networking

import (
	realRand "math/rand"
	"time"
)

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[realRand.Intn(len(letterRunes))]
	}
	return string(b)
}

func generateKey(size int) ([]byte, error) {
	realRand.Seed(time.Now().UnixNano())
	return []byte(randStringRunes(size)), nil
}
