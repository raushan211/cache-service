package main

import (
	pb "cache-service/client/pb/cache"
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCacheServiceClient(conn)
	setRes, err := client.SetValue(context.Background(), &pb.SetRequest{Key: "name", Value: "Raushan"})
	if err != nil {
		log.Println("failed to set value in SetValue: %v", err)
	} else {
		log.Printf("set req response: %v", setRes)
	}

	getRes, err := client.GetValue(context.Background(), &pb.GetRequest{Key: "name"})
	if err != nil {
		log.Println("failed to get value in GetValue: %v", err)
	} else {
		log.Printf("get req value: %v", getRes)
	}

	setUserRes, err := client.SetUser(context.Background(), &pb.SetUserRequest{
		Name:     "Raushan",
		RollNum:  13,
		Class:    "IX",
		Metadata: []byte("hello"),
	})
	if err != nil {
		log.Println("failed to set value in SetValue: %v", err)
	} else {
		log.Printf("set req response: %v", setUserRes)
	}

	getUserRes, err := client.GetUser(context.Background(), &pb.GetUserRequest{
		Name:    "Raushan",
		RollNum: 13,
	})
	if err != nil {
		log.Println("failed to get value in GetValue: %v", err)
	} else {
		log.Printf("get req value: %v", getUserRes)
	}

}

type (
	User struct {
		Name     string `json:"name"`
		Class    string `json:"class"`
		RollNum  int64  `json:"roll_num"`
		Metadata []byte `json:"metadata"`
	}
)
