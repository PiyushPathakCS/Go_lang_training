package main

import (
	"context"
	pb "grpc_project/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GoodServiceClient, names *pb.NameList) {
	log.Printf("Client Streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)

	}

	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Message: name,
		}

		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)

		}
		log.Printf("sent request with name :%s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client Streaming finished ")
	if err != nil {
		log.Printf("Error while receiving %v", err)
	}
	log.Printf("%v", res.Message)
}
