package main

import (
	"context"
	"io"
	"log"

	pb "github.com/Pkra99/grpc_go/proto"
)

func callSayHelloFromServer(client pb.GreetServiceClient, names *pb.NameList) {
	log.Print("Streaming started")
	stream, err := client.SayHelloFromServer(context.Background(), names)
	if err != nil {
		log.Fatalf("Error getting names %v", err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error streaming names %v", err)
		}
		log.Println(message)
	}
	log.Println("Finished streaming")
}
