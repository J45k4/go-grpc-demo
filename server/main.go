package main

import (
	"log"
	"net"

	"github.com/j45k4/go-grpc-demo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
)

type GrpcDemoServer struct {
}

func main() {
	lis, err := net.Listen("tcp", ":7777")
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile("./server.crt", "./server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	GrpcDemoService.RegisterGrpcDemoServiceServer(s, &GrpcDemoServer{})

	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (s *GrpcDemoServer) Hello(ctx context.Context, in *GrpcDemoService.HelloRequest) (*GrpcDemoService.HelloResponse, error) {
	return &GrpcDemoService.HelloResponse{
		ReturnString: "Hello " + in.Name,
	}, nil
}
