package main

import (
	"log"

	pb "github.com/abhishekmaurya0/2fa/proto"
	"github.com/abhishekmaurya0/2fa/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"
var database string = "root:ittfaq1901@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"

func main() {
	var ser Client
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	ser.AuthServiceClient = pb.NewAuthServiceClient(conn)
	ser.DB, err = server.SetupDBConnection(database)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	err = ser.DB.AutoMigrate(&User_response{})
	if err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}
	DoSignupUser(ser)
	DoEnable2fa(ser)
}
