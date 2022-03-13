package main

import (
	"context"
	"fmt"
	"github.com/3110Y/profile/pkg/profileGRPC"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5300", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := profileGRPC.NewProfileServiceClient(conn)
	request := &profileGRPC.ProfileWithoutIdSystemField{
		Email:      "test@test.test",
		Phone:      79062579331,
		Password:   "Password8",
		Surname:    "Surname",
		Name:       "Name",
		Patronymic: "Patronymic",
	}
	response, err := client.Add(context.Background(), request)
	if err != nil {
		grpclog.Fatalf("fail to client: %v", err)
	}
	fmt.Println(response)
}
