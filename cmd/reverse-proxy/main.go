package main

import (
	"context"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"github.com/iam-solutions/user-service/api/v1/pb"
)

func main() {
	port := ":8081"
	userServiceServerEndpoint := "127.0.0.1:8089"

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, userServiceServerEndpoint, opts); err != nil {
		log.Fatalln(err)
	}

	s := &http.Server{
		Addr: port,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			log.Fatalf("Failed to shutdown reverse proxy: %v\n", err.Error())
		}
	}()

	log.Printf("Reverse Proxy is running on %s\n", port)
	if err := s.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start reverse proxy: %v\n", err.Error())
	}
}
