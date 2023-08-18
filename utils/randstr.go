package utils

import (
	"math/rand"
	"time"
)

func RandStr() string {
	words := []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	str := ""
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < 8; i++ {
		index := r.Intn(len(words))
		str += string(words[index])
	}
	return str
}
