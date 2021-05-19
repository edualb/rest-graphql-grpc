package main

import (
	"log"
	"net"

	"github.com/edualb/rest-graphql-grpc/server/grpc/pb"
	"github.com/edualb/rest-graphql-grpc/server/grpc/service"
	"google.golang.org/grpc"
)

func main() {
	log.Printf("Start server on port 8080")

	userService := service.NewUserService()
	flightService := service.NewFlightService()

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, userService)
	pb.RegisterFlightServiceServer(grpcServer, flightService)

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("cannot start server", err)
	}

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start grpc server", err)
	}
}
