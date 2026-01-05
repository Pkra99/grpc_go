package main

import (
	"context"
	"log"
	"time"

	pb "github.com/Pkra99/grpc_go/proto"
)

func callSayHelloFromClient(client pb.GreetServiceClient,  names *pb.NameList) {
	log.Printf("Client streaming started")
	stream, err := client.SayHelloFromClient(context.Background())
	if err != nil {
		log.Fatalf("Error sending name from client %v", err)
	}

	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error streaming names from client %v", err)
		}
		log.Printf("Sent request with the names %s", name)
		time.Sleep(time.Second)
	}
	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err!= nil {
		log.Fatalf("Error while reciving response %v", err)
	}
	log.Printf("%v", res.Messages)
}