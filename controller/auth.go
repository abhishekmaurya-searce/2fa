package controllers

import (
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"time"
)

func GenerateToken(key *ecdsa.PrivateKey, secret string) ([]byte, error) {
	hash := GenerateHash(secret, time.Now())
	sign, err := ecdsa.SignASN1(rand.Reader, key, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to Sign the Hash")
	}
	return sign, nil
}
