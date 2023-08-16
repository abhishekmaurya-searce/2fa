package graph

import (
	pb "github.com/abhishekmaurya0/2fa/proto"
	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Client pb.AuthServiceClient
	DB     *gorm.DB
}
