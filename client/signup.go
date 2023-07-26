package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/abhishekmaurya0/2fa/proto"
)

func doSignUpUser(c pb.AuthServiceClient) {
	res, err := c.SignUpUser(context.Background(), &pb.RegisterUserRequest{
		Name:     "Abhishek Maurya",
		Email:    "abhishek.maurya@searce.com",
		Password: "securepassword",
	})
	if err != nil {
		log.Fatalf("Couldn't register user: %v", err)
	}
	fmt.Print(res)
}
