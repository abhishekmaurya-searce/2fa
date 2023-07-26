package main

import (
	"log"
	"net"

	"github.com/abhishekmaurya0/2fa/models"
	pb "github.com/abhishekmaurya0/2fa/proto"
	"github.com/abhishekmaurya0/2fa/server"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"
var database string = "root:ittfaq1901@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	var ser server.Server
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to Listen: %v", err)
	}
	log.Printf("Listning to: %s", addr)
	ser.DB, err = server.SetupDBConnection(database)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}

	err = ser.DB.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAuthServiceServer(s, &ser)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to Serve: %v", err)
	}

}
