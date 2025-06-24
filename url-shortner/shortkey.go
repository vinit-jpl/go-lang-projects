package main

import (
	"math/rand"
	"time"
)

func GenerateShortKey() string {
	const charSet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	const keyLength = 6

	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	shortKey := make([]byte, keyLength)

	for i := range shortKey {
		shortKey[i] = charSet[seededRand.Intn(len(charSet))]
	}

	return string(shortKey)
}
