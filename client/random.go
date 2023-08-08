package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	// Allowed characters for the local part of the email address
	localPartChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Domain name for the email address
	domain = "example.com"
)

func generateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, length)
	for i := range result {
		result[i] = localPartChars[rand.Intn(len(localPartChars))]
	}
	return string(result)
}

func generateRandomEmail() string {
	localPart := generateRandomString(10)
	email := fmt.Sprintf("%s@%s", localPart, domain)
	return email
}
