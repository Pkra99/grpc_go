package main

import (
	"context"

	pb "github.com/Pkra99/grpc_go/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoPrams) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}
