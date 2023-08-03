package services

import (
	"context"

	pb "github.com/gokul656/ms-arch/payment-service/app/protobufs"
)

type GreetService struct {
	pb.UnimplementedGreetServiceServer
}

func (g *GreetService) SayHello(ctx context.Context, message *pb.GreetRequest) (*pb.GreetRequest, error) {

	resp := pb.GreetRequest{
		Message: "Hello from server!",
	}

	return &resp, nil
}
