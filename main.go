package main

import (
	"context"
	"log"
	"net"

	pb "cache-service/server/pb/cache"
	"cache-service/server/pb/cache/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedCacheServiceServer
	Logic *service.RedisDB
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	logic := service.NewRedisDB(log.Default())
	ser := &server{
		Logic: logic,
	}
	pb.RegisterCacheServiceServer(s, ser)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *server) GetValue(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	val, err := s.Logic.Get(context.Background(), in.Key)
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, err
	}
	return &pb.GetResponse{
		Value: string(val),
	}, nil
}

func (s *server) SetValue(ctx context.Context, in *pb.SetRequest) (*pb.SetResponse, error) {
	log.Printf("Received request: %v", in.ProtoReflect().Descriptor().FullName())
	err := s.Logic.Set(context.Background(), in.Key, []byte(in.Value))
	if err != nil {
		log.Printf("Error: %v", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.SetResponse{
		Message: "Value set successfully",
	}, nil
}
