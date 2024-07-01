package main

import (
	pb "grpc_project/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8000"
)

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not Connect: %v", err)

	}
	defer conn.Close()

	client := pb.NewGoodServiceClient(conn)
	names := &pb.NameList{
		Name: []string{"Akhil", "Alice", "bob"},
	}

	//callSayHelloServerStream(client, names)
	//callSayHelloClientStream(client, names)
	callHelloBidirectionalStream(client, names)
	//callSayHello(client)
}
