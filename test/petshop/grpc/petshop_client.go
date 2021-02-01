package main

import (
	"context"
	"log"
	"time"

	pb "github.com/funa1g/microservice-example/api/petshop"
	"google.golang.org/grpc"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewPetshopClient(conn)

	// Contact the server and print out its response.
	limit := 10
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.GetPetList(ctx, &pb.GetPetListRequest{Limit: int32(limit)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	for _, pet := range r.GetResults() {
		log.Printf("Pet Name: %s", pet.GetName())
	}
}
