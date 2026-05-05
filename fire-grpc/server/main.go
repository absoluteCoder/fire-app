package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "fire-grpc/gen/pb"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFIREServiceServer
}

func (s *server) Ignite(ctx context.Context, req *pb.FireRequest) (*pb.FireResponse, error) {
	msg := fmt.Sprintf("🔥 Fire ignited using %s!", req.Fuel)
	return &pb.FireResponse{Message: msg}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterFIREServiceServer(grpcServer, &server{})

	fmt.Println("🚀 Server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
