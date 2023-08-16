package main

// import (
// 	"context"
// 	"fmt"
// 	"strings"

// 	proto "github.com/abhishekmaurya0/2fa/proto"
// )

// func Enable2fa(s Client) (string, error) {
// 	email := SignUpUser(s)
// 	var user User_response
// 	result := s.DB.First(&user, "email = ?", strings.ToLower(email))
// 	if result.Error != nil {
// 		return "", fmt.Errorf("can't enable 2fa: %s", result.Error)
// 	}
// 	res, err := s.Enable2FA(context.Background(), &proto.LoginUserRequest{
// 		Email:    user.Email,
// 		Password: user.Password,
// 	})
// 	if err != nil {
// 		return "", fmt.Errorf("error in enabling 2fa: %s", err)
// 	}
// 	user.Secret = res.OtpSecret
// 	s.DB.Save(&user)
// 	return email, nil
// }

// func DoEnable2fa(s Client) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Print(Enable2fa(s))
// 	}
// }
