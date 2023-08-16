package controllers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

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
