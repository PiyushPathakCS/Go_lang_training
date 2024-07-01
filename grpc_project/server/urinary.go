package main

import (
	"context"
	pb "grpc_project/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	name := req.Name
	return &pb.HelloResponse{
		Message: "Hello" + name,
	}, nil
}
