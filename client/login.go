package main

import (
	"context"
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
	s.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    user.Email,
		Password: user.Password,
		Otp:      auth.GenerateTOTP(user.Secret, time.Now()),
		Token: ,
	})
}
