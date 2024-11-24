package main

import (
	"fmt"
	"log"
	"net"

	pb "decoder/proto"

	"google.golang.org/grpc"
)

type DecoderServiceServer struct {
	pb.UnimplementedDecoderServiceServer
}

func main() {
	server := grpc.NewServer()
	pb.RegisterDecoderServiceServer(server, &DecoderServiceServer{})

	Port := "50052"

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	log.Println(fmt.Sprintf("Starting gRPC server on port %s...", Port))
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
