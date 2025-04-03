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
	rentAmount, err := decimal.NewFromString(request.RentAmount)
	if err != nil {
		log.Println("Invalid rent amount format:", err)
		return &pb.HoldResponse{Success: false}, err
	}

	pledgeAmount, err := decimal.NewFromString(request.PledgeAmount)
	if err != nil {
		log.Println("Invalid pledge amount format:", err)
		return &pb.HoldResponse{Success: false}, err
	}

	heldFundsID, err := h.paymentUseCase.HoldFunds(ctx, request.RenterId, rentAmount, pledgeAmount)
	if err != nil {
		log.Println("Hold error:", err)
		return &pb.HoldResponse{Success: false}, err
	}

	return &pb.HoldResponse{Success: true, HeldFundsID: heldFundsID}, nil
}

func (h *PaymentHandler) CompleteRent(ctx context.Context, request *pb.CompleteRentRequest) (*pb.CompleteRentResponse, error) {
	err := h.paymentUseCase.CompleteRent(ctx, request.RenterId, request.LandlordId, request.HeldFundsID, request.ToLandlord)
	if err != nil {
		log.Println("Complete rent error:", err)
		return &pb.CompleteRentResponse{Success: false}, err
	}

	return &pb.CompleteRentResponse{Success: true}, nil
}
