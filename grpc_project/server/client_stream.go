package main

import (
	pb "grpc_project/proto"
	"io"
	"log"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GoodService_SayHelloClientStreamingServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Message)
		messages = append(messages, "Hello", req.Message)
	}

}
