package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"context"

	pb "github.com/mnuma/sample-grpc/era/proto"
)

type EraService struct {
}

func (s *EraService) GetCurrentEra(ctx context.Context, message *pb.CurrentEraMessage) (*pb.CurrentEraResponse, error) {
	return &pb.CurrentEraResponse{
		Name: "令和",
	}, nil
}

func main() {
	listenPort, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	service := &EraService{}
	pb.RegisterEraServer(server, service)
	log.Print("start EraService... [:50051]")
	server.Serve(listenPort)
}
