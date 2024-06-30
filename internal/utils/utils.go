package utils

import (
	"fmt"
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwuxzABCDEFGHIJKLMNOPQRSTUVWUXZ0123456789")

func GenerateKey(size int) string {
	b := make([]rune, size)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GetISOTimestamp(t *time.Time) string {
	if t == nil {
		tt := time.Now()
		t = &tt
	}

	return fmt.Sprint(t.UTC().Format(time.RFC3339))
}
