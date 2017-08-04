package main

import (
	"fmt"
	"log"

	"github.com/j45k4/go-grpc-demo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:7777", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := GrpcDemoService.NewGrpcDemoServiceClient(conn)

	r, err := c.Hello(context.Background(), &GrpcDemoService.HelloRequest{
		Name: "Jaska",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r.ReturnString)
}
