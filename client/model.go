package main

import (
	"fmt"

	pb "github.com/abhishekmaurya0/2fa/proto"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Client struct {
	pb.AuthServiceClient
	*gorm.DB
}

type User_response struct {
	Name        string `gorm:"type:varchar(255);not null"`
	Password    string `gorm:"type:varchar(255);not null"`
	Secret      string `gorm:"type:varchar(255)"`
	Email       string `gorm:"primary_key;type:varchar(255)"`
	Private_key string `gorm:"type:varchar(255);not null"`
}

func SetupDBConnection(dsn string) (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("mysql connection err: %v", err)
	}
	return db, nil
}
