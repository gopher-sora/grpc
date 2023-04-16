package main

import (
	"log"
	"net"

	pb "github.com/gopher-sora/grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	list, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}
	log.Printf("listening on %s\n", addr)

	s := grpc.NewServer()

	if err := s.Serve(list); err != nil {
		log.Fatalf("Failed to server: %v\n", err)
	}
}
