package main

import (
	"context"
	"log"
	"time"

	"github.com/gokul656/ms-arch/payment-service/app/protobufs"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func main() {
	log.Println("[BOOT] STARTING gRPC CLIENT...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	ctx = metadata.NewOutgoingContext(ctx, metadata.New(map[string]string{"x-api-key": "gokul656"}))

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	conn, err := grpc.Dial("localhost:3500", opts...)
	if err != nil {
		log.Fatalln("[ERROR]", err.Error())
	}
	defer conn.Close()

	client := protobufs.NewPaymentServiceClient(conn)
	response, err := client.ProcessRequest(ctx, &protobufs.PaymentRequest{
		Uid:           "Gokul656",
		RequestId:     uuid.NewString(),
		Amount:        1000,
		PaymentAction: protobufs.PaymentAction_DEPOSIT,
	})

	if err != nil {
		log.Println("[ERROR]", err.Error())
	}

	greetClient := protobufs.NewGreetServiceClient(conn)
	greetCtx, greetCancel := context.WithTimeout(context.Background(), time.Second)
	greetCtx = metadata.NewOutgoingContext(greetCtx, metadata.New(map[string]string{"x-api-key": "gokul656"}))

	greeting, err := greetClient.SayHello(greetCtx, &protobufs.GreetRequest{
		Message: "Hello Server!",
	})

	if err != nil {
		log.Println("[ERROR]", err.Error())
	}

	log.Println("[RESPONSE]", response)
	log.Println("[RESPONSE]", greeting)

	defer cancel()
	defer greetCancel()
}
