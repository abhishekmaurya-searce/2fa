package controllers

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

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
