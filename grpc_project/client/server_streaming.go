package main

import (
	"context"
	pb "grpc_project/proto"
	"io"
	"log"
)

func callSayHelloServerStream(client pb.GoodServiceClient, names *pb.NameList) {
	log.Printf("Streaming has started ")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalf("Could not Steam: %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while Streaming %v", err)
		}
		log.Println(message)
	}
	log.Fatalf("Streaming Finished")
}
