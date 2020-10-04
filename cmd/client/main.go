package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"

	"github.com/iam-solutions/user-service/api/v1/pb"
)

func main() {
	conn, err := grpc.Dial("localhost:8089", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	userServiceClient := pb.NewUserServiceClient(conn)

	stream, err := userServiceClient.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		if err == io.EOF {
			return
		}

		user := res.GetUser()

		fmt.Println(user)
	}
}
