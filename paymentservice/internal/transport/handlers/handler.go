package handler

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"log"
	pb "paymentservice/api/payment"
	"paymentservice/internal/domain/usecases"
)

type PaymentHandler struct {
	pb.UnimplementedPaymentServiceServer
	paymentUseCase usecases.PaymentUseCaseImpl
}

func NewPaymentHandler(paymentUseCase usecases.PaymentUseCaseImpl) *PaymentHandler {
	return &PaymentHandler{paymentUseCase: paymentUseCase}
}

func (h *PaymentHandler) Deposit(ctx context.Context, request *pb.DepositRequest) (*pb.DepositResponse, error) {
	if request.Amount == nil {
		return &pb.DepositResponse{Success: false}, errors.New("amount is required")
	}

	amountStr := request.Amount.GetValue()
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		log.Println("error: invalid amount format:", err)
		return &pb.DepositResponse{Success: false}, err
	}

	err = h.paymentUseCase.Deposit(ctx, request.UserId, amount)
	if err != nil {
		log.Println("deposit error:", err)
		return &pb.DepositResponse{Success: false}, err
	}

	return &pb.DepositResponse{Success: true}, nil
}

func (h *PaymentHandler) Withdraw(ctx context.Context, request *pb.WithDrawRequest) (*pb.WithDrawResponse, error) {
	if request.Amount == nil {
		return &pb.WithDrawResponse{Success: false}, errors.New("amount is required")
	}

	amountStr := request.Amount.GetValue()
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		log.Println("error: invalid amount format:", err)
		return &pb.WithDrawResponse{Success: false}, err
	}

	err = h.paymentUseCase.Withdraw(ctx, request.UserId, amount)
	if err != nil {
		log.Println("withdraw error:", err)
		return &pb.WithDrawResponse{Success: false}, err
	}

	return &pb.WithDrawResponse{Success: true}, nil
}

func (h *PaymentHandler) Hold(ctx context.Context, request *pb.HoldRequest) (*pb.HoldResponse, error) {
	if request.Amount == nil || request.RentId == 0 {
		return &pb.HoldResponse{Success: false}, errors.New("amount and rent_id are required")
	}

	amountStr := request.Amount.GetValue()
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		log.Println("Invalid amount format:", err)
		return &pb.HoldResponse{Success: false}, err
	}

	err = h.paymentUseCase.HoldFunds(ctx, request.UserId, amount, request.RentId)
	if err != nil {
		log.Println("Hold error:", err)
		return &pb.HoldResponse{Success: false}, err
	}

	return &pb.HoldResponse{Success: true}, nil
}

func (h *PaymentHandler) Pay(ctx context.Context, request *pb.PayRequest) (*pb.PayResponse, error) {
	if request.RentAmount == nil || request.PledgeAmount == nil {
		return &pb.PayResponse{Success: false}, errors.New("both rent_amount and pledge_amount are required")
	}

	rentStr := request.RentAmount.GetValue()
	pledgeStr := request.PledgeAmount.GetValue()

	rentAmount, err := decimal.NewFromString(rentStr)
	if err != nil {
		log.Println("error: invalid rent amount format:", err)
		return &pb.PayResponse{Success: false}, err
	}

	pledgeAmount, err := decimal.NewFromString(pledgeStr)
	if err != nil {
		log.Println("error: invalid pledge amount format:", err)
		return &pb.PayResponse{Success: false}, err
	}

	// Добавляем RentId в вызов PayRent
	err = h.paymentUseCase.PayRent(ctx, request.UserId, request.LandlordId, request.RentId, rentAmount, pledgeAmount)
	if err != nil {
		log.Println("pay error:", err)
		return &pb.PayResponse{Success: false}, err
	}

	return &pb.PayResponse{Success: true}, nil
}

func (h *PaymentHandler) CompleteRent(ctx context.Context, request *pb.CompleteRentRequest) (*pb.CompleteRentResponse, error) {
	if request.RentId == 0 {
		return &pb.CompleteRentResponse{Success: false}, errors.New("rent_id is required")
	}

	err := h.paymentUseCase.CompleteRent(ctx, request.RentId, request.ToLandlord)
	if err != nil {
		log.Println("complete rent error:", err)
		return &pb.CompleteRentResponse{Success: false}, err
	}

	return &pb.CompleteRentResponse{Success: true}, nil
}
