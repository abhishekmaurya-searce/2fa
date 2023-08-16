package graph

import (
	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	pb "github.com/abhishekmaurya0/2fa/proto"
	"github.com/abhishekmaurya0/2fa/server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"
var database string = "root:ittfaq1901@tcp(127.0.0.1:3306)/auth?charset=utf8mb4&parseTime=True&loc=Local"

func SetUpClient() {
	var ser Resolver
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()
	ser.Client = pb.NewAuthServiceClient(conn)
	ser.DB, err = server.SetupDBConnection(database)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	err = ser.DB.AutoMigrate(&User_response{})
	if err != nil {
		log.Fatalf("Failed to migrate schema: %v", err)
	}
	serve := handler.NewDefaultServer(
		NewExecutableSchema(
			Config{Resolvers: &mutationResolver{Resolver: &ser}},
		),
	)
	serve.AddTransport(transport.POST{})

	http.Handle("/auth", serve)
	http.ListenAndServe(":8080", nil)

}
