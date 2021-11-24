package main

import (
	"context"
	"log"
	"math"
	"net"

	"hello"

	"github.com/kelseyhightower/envconfig"
	"google.golang.org/grpc"
)

type Config struct {
	Host string `envconfig:"HOST" required:"true" default:":9090"`
}

func getConfig() Config {
	config := Config{}

	err := envconfig.Process("HELLOGRPC", &config)
	if err != nil {
		log.Fatal(err)
	}

	return config
}

func main() {
	config := getConfig()
	lis, err := net.Listen("tcp", config.Host)

	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	hello.RegisterHelloServiceServer(server, &helloServiceServer{})

	log.Printf("listening for gRPC requests on %s", config.Host)

	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}

type helloServiceServer struct {
	hello.HelloServiceServer
}

func (s *helloServiceServer) GetMessage(ctx context.Context, req *hello.Nothing) (*hello.HelloResponse, error) {
	log.Printf("gRPC request received")

	return &hello.HelloResponse{
		Message: "Hello, gRPC!",
	}, nil
}

func (s *helloServiceServer) GenerateLoad(ctx context.Context, req *hello.Nothing) (*hello.Nothing, error) {
	x := 0.0001

	for i := 0; i <= 1000000; i++ {
		x = x + math.Sqrt(x)
	}

	return &hello.Nothing{}, nil
}
