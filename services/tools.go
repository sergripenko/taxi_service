package services

import (
	"crypto/rand"
)

const (
	letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	lenKey  = 2
)

// Random generate string
func GetRandomString() string {
	var bytes = make([]byte, lenKey)
	rand.Read(bytes)

	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return string(bytes)
}
