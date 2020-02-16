package main

import (
	pb "GoTest/grpc/example2/proto"
	"context"
	"google.golang.org/grpc"
	"log"
	"os"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the service.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewToUpperClient(conn)

	// Contact the service and print out its response.
	name := "hello world"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	r, err := c.Upper(context.Background(), &pb.UpperRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.Message)
}
