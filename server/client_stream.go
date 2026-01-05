package main

import (
	"io"
	"log"

	pb "github.com/Pkra99/grpc_go/proto"
)

func (s *helloServer) SayHelloFromClient(stream pb.GreetService_SayHelloFromClientServer) error {
	var messages []string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Messages: messages})
		}
		if err != nil {
			log.Fatalf("Error streaming name %v", err)
		}
		log.Printf("Requested names: %v", req.Name)
		messages = append(messages, "Hello", req.Name)
	}
}
