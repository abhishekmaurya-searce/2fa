package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"strings"

// 	pb "github.com/abhishekmaurya0/2fa/proto"
// )

// func SignUpUser(s Client) string {
// 	name := generateRandomString(15)
// 	email := strings.ToLower(generateRandomEmail())
// 	pass := generateRandomString(20)
// 	res, err := s.SignUpUser(context.Background(), &pb.RegisterUserRequest{
// 		Name:     name,
// 		Email:    email,
// 		Password: pass,
// 	})
// 	if err != nil {
// 		log.Fatalf("Couldn't register user: %v", err)
// 		return ""
// 	}
// 	user := &User_response{
// 		Name:        name,
// 		Email:       email,
// 		Password:    pass,
// 		Private_key: res.PrivateKey,
// 	}
// 	s.DB.Create(&user)
// 	return email
// }

// func DoSignupUser(s Client) {

// 	for i := 0; i < 10; i++ {
// 		fmt.Println(SignUpUser(s))
// 	}
// }
