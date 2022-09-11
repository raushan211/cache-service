package main

import (
	"context"
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

func generateResponse(value []byte, err error) (*pb.ServerResponse, error) {
	if err != nil {
		return &pb.ServerResponse{Success: false, Value: string(value), Error: err.Error()}, nil
	}
	return &pb.ServerResponse{Success: true, Value: string(value), Error: ""}, nil
}
