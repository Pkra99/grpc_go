package main

import (
	"log"

	pb "github.com/Pkra99/grpc_go/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const port = ":8080"

func main() {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Error establishing connection to server %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)
	callSayHello(client)
}
