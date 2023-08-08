package controllers

import (
	"crypto/ecdsa"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base32"
	"encoding/pem"
	"fmt"
	"math"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
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
func GenerateTOTP(secret string, currentTime time.Time) string {
	hmacValue := GenerateHash(secret, currentTime)
	offset := int(hmacValue[19] & 0x0f)
	binary := ((int(hmacValue[offset]) & 0x7f) << 24) |
		((int(hmacValue[offset+1] & 0xff)) << 16) |
		((int(hmacValue[offset+2] & 0xff)) << 8) |
		(int(hmacValue[offset+3] & 0xff))

	digits := 6
	mod := int(math.Pow10(digits))
	otp := binary % mod

	code := fmt.Sprintf("%0*d", digits, otp)

	return code
}

// Verify a TOTP code
func ValidateTOTP(code, secret string, currentTime time.Time) bool {
	generatedCode := GenerateTOTP(secret, currentTime)

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

func GenerateKeys(key *ecdsa.PrivateKey) (string, string) {
	private, _ := x509.MarshalECPrivateKey(key)
	privateKeyPem := &pem.Block{
		Type:  "EC PRIVATE KEY",
		Bytes: private,
	}
	privateKeyBase64 := pem.EncodeToMemory(privateKeyPem)

	public, _ := x509.MarshalPKIXPublicKey(key.PublicKey)
	publicKeyPem := &pem.Block{
		Type:  "EC PUBLIC KEY",
		Bytes: public,
	}
	publicKeyBase64 := pem.EncodeToMemory(publicKeyPem)

	return string(privateKeyBase64), string(publicKeyBase64)
}

func PrivateKeyFromPEM(privateKeyPEM string) (*ecdsa.PrivateKey, error) {
	privateKeyPem, _ := pem.Decode([]byte(privateKeyPEM))
	if privateKeyPem == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	if privateKeyPem.Type != "EC PRIVATE KEY" {
		return nil, fmt.Errorf("invalid PEM block type, expected 'EC PRIVATE KEY'")
	}

	privateKey, err := x509.ParseECPrivateKey(privateKeyPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA private key: %w", err)
	}

	return privateKey, nil
}

func PublicKeyFromPEM(publicKeyPEM string) (*ecdsa.PublicKey, error) {
	publicKeyPem, _ := pem.Decode([]byte(publicKeyPEM))
	if publicKeyPem == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing public key")
	}
	fmt.Println(publicKeyPem.Type)
	if publicKeyPem.Type != "EC PUBLIC KEY" {
		return nil, fmt.Errorf("invalid PEM block type, expected 'EC PUBLIC KEY'")
	}

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse ECDSA public key: %w", err)
	}

	ecdsaPublicKey, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("failed to cast public key to ECDSA public key")
	}

	return ecdsaPublicKey, nil
}

func GenerateToken(key *ecdsa.PrivateKey, secret string) ([]byte, error) {
	hash := GenerateHash(secret, time.Now())
	sign, err := ecdsa.SignASN1(rand.Reader, key, hash)
	if err != nil {
		return nil, fmt.Errorf("failed to Sign the Hash")
	}
	return sign, nil
}
