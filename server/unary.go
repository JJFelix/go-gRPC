package main

import (
	"context"

	pb "github.com/JJFelix/go-gRPC/proto"
)

// works like a regular API

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam)(*pb.HelloResponse, error){
	return &pb.HelloResponse{
		Message: "Hello",
	}, nil
}