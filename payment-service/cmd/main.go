package main

import (
	"log"
	"net"

	"github.com/gokul656/ms-arch/payment-service/app/interceptors"
	pb "github.com/gokul656/ms-arch/payment-service/app/protobufs"
	"github.com/gokul656/ms-arch/payment-service/app/services"
	"google.golang.org/grpc"
)

func main() {
	log.Println("[BOOT] STARTING PAYMENT-SERVICE...")

	listen, err := net.Listen("tcp", "localhost:3500")
	if err != nil {
		log.Fatalln("[ERR]", err.Error())
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptors.ValidateToken),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterPaymentServiceServer(server, &services.PaymentService{})
	pb.RegisterGreetServiceServer(server, &services.GreetService{})

	log.Println("[BOOT] SERVER LISTENING ON :", listen.Addr())

	if err := server.Serve(listen); err != nil {
		log.Fatalln("[ERROR]", err.Error())
	}
}
