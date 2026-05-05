package main

import (
	"context"
	"log"
	"time"

	pb "fire-grpc/gen/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFIREServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Ignite(ctx, &pb.FireRequest{Fuel: "wood"})
	if err != nil {
		log.Fatalf("could not ignite: %v", err)
	}

	log.Printf("Response: %s", res.Message)
}
