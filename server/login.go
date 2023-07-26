package server

import (
	"context"
	"strings"
	"time"

	auth "github.com/abhishekmaurya0/2fa/controller"
	"github.com/abhishekmaurya0/2fa/models"
	pb "github.com/abhishekmaurya0/2fa/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.UserResponse, error) {
	var user models.User
	result := s.DB.First(&user, "email = ?", strings.ToLower(req.Email))
	if result.Error != nil {
		return nil, status.Errorf(codes.NotFound, "User not found")
	}

	if !auth.ValidatePass(req.Password, []byte(user.Password)) {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid password")
	}

	userResponse := &pb.UserResponse{
		Name:       user.Name,
		Email:      user.Email,
		OtpEnabled: user.Otp_enabled,
	}

	if user.Otp_enabled {
		flag := auth.ValidateTOTP(req.Otp, user.Otp_secret, time.Now())
		if !flag {
			return nil, status.Errorf(codes.InvalidArgument, "OTP verification unsuccessful")
		}
	}

	return userResponse, nil
}
