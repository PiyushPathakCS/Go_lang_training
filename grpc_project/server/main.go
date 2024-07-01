package main

import (
	pb "grpc_project/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":8000"
)

type helloServer struct {
	pb.GoodServiceServer
}

func main() {

	lis, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Unable to start server %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGoodServiceServer(grpcServer, &helloServer{})
	log.Printf("Server started at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Unable to start %v", err)
	}

}
