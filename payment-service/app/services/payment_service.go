package services

import (
	"context"
	"errors"

	pb "github.com/gokul656/ms-arch/payment-service/app/protobufs"
	"github.com/google/uuid"
)

type PaymentService struct {
	pb.UnimplementedPaymentServiceServer
}

func (p *PaymentService) ProcessRequest(ctx context.Context, request *pb.PaymentRequest) (*pb.PaymentRequest, error) {

	if request.Uid != "Gokul656" {
		return nil, errors.New("unable to validate")
	}

	request.Amount += 100
	request.Uid = uuid.NewString()

	return request, nil
}
