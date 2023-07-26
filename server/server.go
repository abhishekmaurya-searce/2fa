package server

import (
	"fmt"

	pb "github.com/abhishekmaurya0/2fa/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Server struct {
	pb.AuthServiceServer
	*gorm.DB
}

func SetupDBConnection(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("mysql connection err: %v", err)
	}
	return db, nil
}
