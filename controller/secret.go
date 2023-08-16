package controllers

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"math"
	"strings"
	"time"
)

func GenerateSecretKey() string {

	secretBytes := make([]byte, 10)
	_, _ = rand.Read(secretBytes)

	secret := base32.StdEncoding.EncodeToString(secretBytes)

	secret = strings.TrimRight(secret, "=")

	return secret
}

func GenerateHash(secret string, currentTime time.Time) []byte {

	secretBytes, _ := base32.StdEncoding.DecodeString(secret)

	counter := uint64(math.Floor(float64(currentTime.Unix()) / 30))

	counterBytes := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		counterBytes[i] = byte(counter & 0xff)
		counter >>= 8
	}

	hash := hmac.New(sha1.New, secretBytes)
	_, _ = hash.Write(counterBytes)
	hmacValue := hash.Sum(nil)
	return hmacValue
}
