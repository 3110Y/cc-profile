package main

import (
	"fmt"
	"github.com/3110Y/profile/internal/infrastructure/di"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"net"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%s", os.Getenv("GRPC_PORT")))
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	initializeDI, err := di.InitializeDI()
	if err != nil {
		grpclog.Fatalf("failed", err)
	}
	profileGRPC.RegisterProfileServiceServer(grpcServer, initializeDI.ProfileRPC)
	err = grpcServer.Serve(listener)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}
}
