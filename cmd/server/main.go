package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	"github.com/iam-solutions/user-service/api/v1/pb"
	"github.com/iam-solutions/user-service/internal/app/service"
)

var (
	name, serverAddr, gatewayAddr string
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file")
	}

	name = os.Getenv("SERVICE_NAME")
	if name == "" {
		name = "User Service"
	}

	serverAddr = os.Getenv("SERVICE_PORT")
	if serverAddr == "" {
		serverAddr = ":8088"
	} else {
		serverAddr = ":" + serverAddr
	}

	gatewayAddr = os.Getenv("GRPC_GATEWAY_PORT")
	if gatewayAddr == "" {
		gatewayAddr = ":8087"
	} else {
		gatewayAddr = ":" + gatewayAddr
	}
}

func runGRPCServer(ctx context.Context) error {
	lis, err := net.Listen("tcp", serverAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	gRPCServer := grpc.NewServer()
	userServiceServer := service.NewUserServiceServer()
	pb.RegisterUserServiceServer(gRPCServer, userServiceServer)

	go func() {
		defer gRPCServer.GracefulStop()
		<-ctx.Done()
	}()

	return gRPCServer.Serve(lis)
}

func runGRPCGateway(ctx context.Context, opts ...runtime.ServeMuxOption) error {
	mux := runtime.NewServeMux(opts...)
	gRPCOpts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, "localhost"+serverAddr, gRPCOpts); err != nil {
		fmt.Printf("Failed to register endpoint server: %v\n", err.Error())
	}

	s := &http.Server{
		Addr:    gatewayAddr,
		Handler: mux,
	}

	go func() {
		<-ctx.Done()
		if err := s.Shutdown(context.Background()); err != nil {
			fmt.Printf("Failed to shutdown gRPC Gateway: %v\n", err.Error())
		}
	}()

	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		fmt.Printf("Failed to listen and serve: %s\n", err.Error())
		return err
	}

	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("Received %v. Shutting down...\n", sig)
		os.Exit(0)
	}()

	go func() {
		log.Printf("Starting gRPC Gateway on %s\n", gatewayAddr)
		if err := runGRPCGateway(ctx); err != nil {
			log.Fatalf("Failed to start gRPC Gateway: %s\n", err.Error())
		}
	}()

	log.Printf("Starting gRPC Server on %s\n", serverAddr)
	if err := runGRPCServer(ctx); err != nil {
		log.Fatalf("Failed to start gRPC Server: %s\n", err.Error())
	}
}
