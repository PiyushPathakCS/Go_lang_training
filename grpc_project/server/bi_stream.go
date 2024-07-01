package main

import (
	pb "grpc_project/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GoodService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Message)
		res := &pb.HelloResponse{
			Message: "Hello" + req.Message,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
