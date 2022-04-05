package main

import (
	"context"
	"flag"
	"fmt"
	"grpc_demo/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

type server struct {
	hellos []*pb.HelloRequest
}

func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReplqy, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReplqy{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterAnotherServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
