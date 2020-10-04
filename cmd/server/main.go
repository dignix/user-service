package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/iam-solutions/user-service/api/v1/pb"
	"github.com/iam-solutions/user-service/internal/app/service"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8089")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	userServiceServer := service.NewUserServiceServer()

	pb.RegisterUserServiceServer(grpcServer, userServiceServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
