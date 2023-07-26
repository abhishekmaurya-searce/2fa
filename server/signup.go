package server

import (
	"context"
	"strings"

	auth "github.com/abhishekmaurya0/2fa/controller"
	"github.com/abhishekmaurya0/2fa/models"
	pb "github.com/abhishekmaurya0/2fa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) SignUpUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.UserResponse, error) {
	pass, err := auth.GeneratePassword(req.Password)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Wrong Password")
	}
	req.Password = string(pass)

	newUser := models.User{
		Name:        req.Name,
		Email:       strings.ToLower(req.Email),
		Password:    req.Password,
		Otp_enabled: false,
	}

	result := s.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return nil, status.Errorf(codes.AlreadyExists, "Email already exists, please use another email address")
	} else if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create user")
	}
	var userres pb.UserResponse
	userres.Email = newUser.Email
	userres.OtpEnabled = false
	userres.OtpSecret = "nil"
	return &userres, nil
}