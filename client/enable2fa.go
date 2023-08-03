package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/abhishekmaurya0/2fa/proto"
)

func DoEnable2fa(c pb.AuthServiceClient) {
	res, err := c.Enable2FA(context.Background(), &pb.LoginUserRequest{
		Email:    "abhishek.maurya@searce.com",
		Password: "securepassword",
		Otp:      "nil",
	})
	if err != nil {
		log.Fatalf("Couldn't register user: %v", err)
	}
	fmt.Println(res)
}
