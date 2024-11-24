package main

import (
	"context"
	"fmt"
	"log"

	pb "decoder/proto"
)

func DoDecoder(c pb.DecoderServiceClient, txt string) {
	res, err := c.Decode(context.Background(), &pb.DecoderRequest{EncodeText: txt})
	if err != nil {
		log.Fatalf("failed to call Decode: %v", err)
	}

	fmt.Printf("\noutput: %s\n", res.GetData())
}
