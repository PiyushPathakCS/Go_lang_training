package main

import (
	"context"
	pb "grpc_project/proto"
	"io"
	"log"
	"time"
)

func callHelloBidirectionalStream(client pb.GoodServiceClient, names *pb.NameList) {
	log.Printf("Bidirectional Streaming Started")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send names: %v", err)
	}
	waitc := make(chan struct{})

	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		}
		close(waitc)
	}()

	for _, name := range names.Name {
		req := &pb.HelloRequest{
			Message: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending %v", err)
		}
		time.Sleep(2 * time.Second)

	}
	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional streaming finished")
}
