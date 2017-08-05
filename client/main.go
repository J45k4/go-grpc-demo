package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/j45k4/go-grpc-demo/service"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {

	peerCert, err := tls.LoadX509KeyPair("client.crt", "client.key")
	if err != nil {

	}
	caCert, err := ioutil.ReadFile("ca.crt")
	if err != nil {
		// log.Error("read ca cert file error:%v", err)
		// return nil, err
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	ta := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{peerCert},
		RootCAs:      caCertPool,
	})

	conn, err := grpc.Dial("192.168.99.100:7777", grpc.WithTransportCredentials(ta))

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
