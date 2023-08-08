package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"strings"
	"time"

	auth "github.com/abhishekmaurya0/2fa/controller"
	pb "github.com/abhishekmaurya0/2fa/proto"
)

func Login(s Client) error {

	email, err := Enable2fa(s)
	if err != nil {
		return fmt.Errorf("2FA: %s", err)
	}
	var user User_response
	result := s.DB.First(&user, "email = ?", strings.ToLower(email))
	if result.Error != nil {
		return fmt.Errorf("Error in getting data from client DB")
	}
	private_key,err := auth.PrivateKeyFromPEM(user.Private_key)
	hash := auth.GenerateHash(user.Secret,time.Now())
	ecdsa.SignASN1(rand.Reader,private_key,)
	s.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    user.Email,
		Password: user.Password,
		Otp:      auth.GenerateTOTP(user.Secret, time.Now()),
		Token: ,
	})
}
