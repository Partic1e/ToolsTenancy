package handler

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"log"
	pb "paymentservice/api/payment"
	"paymentservice/internal/domain/models"
	"paymentservice/internal/domain/usecases"
	"time"
)

type DepositHandler struct {
	pb.UnimplementedPaymentServer
	useCase *usecases.DepositUseCaseImpl
}

func NewDepositHandler(useCase *usecases.DepositUseCaseImpl) *DepositHandler {
	return &DepositHandler{useCase: useCase}
}

func (h *DepositHandler) Handle(ctx context.Context, request *pb.DepositRequest) (*pb.DepositResponse, error) {
	if request.Amount == nil {
		return &pb.DepositResponse{Success: false}, errors.New("error")
	}
	amountStr := request.Amount.GetValue()

	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		log.Println("", err)
		return &pb.DepositResponse{Success: false}, err
	}

	payment := models.Payment{
		UserId:    uint64(request.UserId),
		Amount:    amount,
		Type:      "deposit",
		CreatedAt: time.Now().Unix(),
	}

	err = h.useCase.Invoke(payment)
	if err != nil {
		log.Println("", err)
		return &pb.DepositResponse{Success: false}, err
	}

	return &pb.DepositResponse{Success: true}, nil
}
