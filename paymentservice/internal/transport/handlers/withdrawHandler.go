package handler

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	pb "paymentservice/api/payment"
	"paymentservice/internal/domain/models"
	"paymentservice/internal/domain/usecases"
	"time"
)

type WithdrawHandler struct {
	pb.UnimplementedPaymentServer
	useCase *usecases.WithdrawUseCaseImpl
}

func NewWithdrawHandler(useCase *usecases.WithdrawUseCaseImpl) *WithdrawHandler {
	return &WithdrawHandler{useCase: useCase}
}

func (h *WithdrawHandler) Handle(ctx context.Context, request *pb.WithDrawRequest) (*pb.WithDrawResponse, error) {
	if request.Amount == nil {
		return &pb.WithDrawResponse{Success: false}, errors.New("amount is not required")
	}
	amountStr := request.Amount.GetValue()

	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		return &pb.WithDrawResponse{Success: false}, err
	}

	payment := models.Payment{
		UserId:    int64(request.UserId),
		Amount:    amount,
		Type:      "withdraw",
		CreatedAt: time.Now(),
	}

	err = h.useCase.Invoke(payment)
	if err != nil {
		return &pb.WithDrawResponse{Success: false}, err
	}

	return &pb.WithDrawResponse{Success: true}, nil
}
