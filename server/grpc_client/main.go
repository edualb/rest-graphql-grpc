package main

import (
	"context"
	"log"

	"github.com/edualb/rest-graphql-grpc/server/grpc/pb"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal("cannot dial server: %w", err)
	}

	userClient := pb.NewUserServiceClient(conn)
	flightClient := pb.NewFlightServiceClient(conn)

	userClient.CreateUser(context.Background(), &pb.CreateUserRequest{
		User: &pb.User{
			Name: "Eduardo Silva",
		},
	})

	flightClient.CreateFlight(context.Background(), &pb.CreateFlightRequest{
		Flight: &pb.Flight{
			From: "Rio de Janeiro",
			To:   "Recife",
		},
	})

}
