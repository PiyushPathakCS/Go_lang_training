package main

import (
	"context"
	pb "grpc_project/proto"
	"log"
	"time"
)

func callSayHello(client pb.GoodServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{
		Name: "Piyush",
	})
	if err != nil {
		log.Fatalf("Could not work %v", err)
	}
	log.Printf("%s", res.Message)
}
