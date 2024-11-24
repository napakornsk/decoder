package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	pb "decoder/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	client := pb.NewDecoderServiceClient(conn)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter txt for decoder: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("failed to read input: %v", err)
	}

	input = strings.TrimSpace(input)

	DoDecoder(client, input)

}
