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

func (s *Server) Enable2FA(ctx context.Context, req *pb.LoginUserRequest) (*pb.UserResponse, error) {
	var user models.User
	result := s.DB.First(&user, "email = ?", strings.ToLower(req.Email))
	if result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	if !auth.ValidatePass(req.Password, []byte(user.Password)) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid Password")
	}
	user.Otp_enabled = true
	user.Otp_secret = auth.GenerateSecretKey()
	result = s.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update user")
	}
	userResponse := &pb.UserResponse{
		Name:       user.Name,
		Email:      user.Email,
		OtpEnabled: user.Otp_enabled,
		OtpSecret:  user.Otp_secret,
	}

	return userResponse, nil
}
