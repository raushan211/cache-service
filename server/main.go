package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"

	"cache-service/server/database"
	pb "cache-service/server/pb/cache"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedCacheServiceServer
	db database.Database
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	databaseImplementation := os.Args[1]
	db, err := database.Factory(databaseImplementation)
	if err != nil {
		panic(err)
	}
	ser := &server{
		db: db,
	}
	pb.RegisterCacheServiceServer(s, ser)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) SetValue(ctx context.Context, in *pb.SetRequest) (*pb.ServerResponse, error) {
	value, err := s.db.Set(in.GetKey(), in.GetValue())
	return generateResponse(value, err)
}
func (s *server) GetValue(ctx context.Context, in *pb.GetRequest) (*pb.ServerResponse, error) {
	value, err := s.db.Get(in.GetKey())
	return generateResponse(value, err)
}

func (s *server) SetUser(ctx context.Context, in *pb.SetUserRequest) (*pb.ServerResponse, error) {
	key := fmt.Sprint(in.Name, ":", in.RollNum)
	user := User{
		Name:     in.Name,
		Class:    in.Class,
		RollNum:  in.RollNum,
		Metadata: in.Metadata,
	}
	value, err := s.db.Set(key, user)
	return generateResponse(value, err)

}

func (s *server) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	key := fmt.Sprint(in.Name, ":", in.RollNum)
	value, err := s.db.Get(key)
	return generateUserResponse(value, err)
}

func generateResponse(value []byte, err error) (*pb.ServerResponse, error) {
	if err != nil {
		return &pb.ServerResponse{Success: false, Value: string(value), Error: err.Error()}, nil
	}
	return &pb.ServerResponse{Success: true, Value: string(value), Error: ""}, nil
}
func generateUserResponse(value []byte, err error) (*pb.GetUserResponse, error) {
	user := User{}
	json.Unmarshal(value, &user)
	if err != nil {
		return &pb.GetUserResponse{Success: false, Error: err.Error()}, nil
	}
	return &pb.GetUserResponse{
		Success:  true,
		Name:     user.Name,
		RollNum:  user.RollNum,
		Class:    user.Class,
		Metadata: user.Metadata,
		Error:    "",
	}, nil
}

type (
	User struct {
		Name     string `json:"name"`
		Class    string `json:"class"`
		RollNum  int64  `json:"roll_num"`
		Metadata []byte `json:"metadata"`
	}
)

func (i User) MarshalBinary() (data []byte, err error) {
	bytes, err := json.Marshal(i)
	return bytes, err
}
