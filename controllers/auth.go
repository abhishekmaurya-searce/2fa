package controllers

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base32"
	"fmt"
	"math"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Generate a secret key for TOTP
func generateSecretKey() string {
	// Generate a random byte array for the secret
	secretBytes := make([]byte, 10)
	_, _ = rand.Read(secretBytes)

	// Encode the secret as base32
	secret := base32.StdEncoding.EncodeToString(secretBytes)

	// Remove any padding characters
	secret = strings.TrimRight(secret, "=")

	return secret
}

// Generate a TOTP code
func generateTOTP(secret string, currentTime time.Time) string {
	// Convert the secret from base32 to bytes
	secretBytes, _ := base32.StdEncoding.DecodeString(secret)

	// Calculate the counter value
	counter := uint64(math.Floor(float64(currentTime.Unix()) / 30))

	// Convert the counter to bytes
	counterBytes := make([]byte, 8)
	for i := 7; i >= 0; i-- {
		counterBytes[i] = byte(counter & 0xff)
		counter >>= 8
	}

	// Calculate the HMAC-SHA1 value
	hash := hmac.New(sha1.New, secretBytes)
	_, _ = hash.Write(counterBytes)
	hmacValue := hash.Sum(nil)

	// Dynamic Truncation
	offset := int(hmacValue[19] & 0x0f)
	binary := ((int(hmacValue[offset]) & 0x7f) << 24) |
		((int(hmacValue[offset+1] & 0xff)) << 16) |
		((int(hmacValue[offset+2] & 0xff)) << 8) |
		(int(hmacValue[offset+3] & 0xff))

	// Generate the TOTP code
	digits := 6
	mod := int(math.Pow10(digits))
	otp := binary % mod

	// Pad the code with leading zeros
	code := fmt.Sprintf("%0*d", digits, otp)

	return code
}

// Verify a TOTP code
func validateTOTP(code, secret string, currentTime time.Time) bool {
	// Generate the TOTP code using the current time
	generatedCode := generateTOTP(secret, currentTime)

	// Compare the generated code with the provided code
	return code == generatedCode
}

func GeneratePassword(pass string) ([]byte, error) {
	hash, err := bcrypt.GenerateFromPassword(([]byte(pass)), bcrypt.DefaultCost)
	if err != nil {
		return hash, err
	}
	return hash, nil
}

func ValidatePass(password string, hash []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		fmt.Println(err)
	}
	return err == nil
}
