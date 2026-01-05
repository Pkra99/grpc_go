package main

import (
	"log"
	"time"

	pb "github.com/Pkra99/grpc_go/proto"
)

func (s *helloServer) SayHelloFromServer(req *pb.NameList, stream pb.GreetService_SayHelloFromServerServer) error {
	log.Printf("Names: %s", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}

		if err := stream.Send(res); err != nil {
			log.Fatalf("Error getting the names %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	return nil

}
