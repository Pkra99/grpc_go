package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Pkra99/grpc_go/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoPrams{})
	if err != nil {
		log.Fatalf("Error getting message %v", err)
	}
	log.Printf("%s", res.Message)
}
