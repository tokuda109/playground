package main

import (
	"log"
	"net"

	pb "github.com/tokuda109/playground/grpc/backend/proto"

	grpc "google.golang.org/grpc"
	reflection "google.golang.org/grpc/reflection"
)

const (
	port = ":50555"
)

type server struct{}

func main() {
	// lis, err := net.Listen("tcp", port)
	// if err != nil {
	// 	log.Fatalf("failed to listen: %v", err)
	// }
	// s := grpc.NewServer()
	// pb.RegisterGreeterServer(s, &server{})
	// reflection.Register(s)
	// if err := s.Serve(lis); err != nil {
	// 	log.Fatalf("failed to serve: %v", err)
	// }
}
