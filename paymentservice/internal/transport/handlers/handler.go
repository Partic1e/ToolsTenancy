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

type PaymentHandler struct {
	pb.UnimplementedPaymentServer
	depositUseCase  *usecases.DepositUseCaseImpl
	withdrawUseCase *usecases.WithdrawUseCaseImpl
	holdUseCase     *usecases.HoldUseCaseImpl
	payUseCase      *usecases.PayUseCaseImpl
}

func NewPaymentHandler(
	depositUseCase *usecases.DepositUseCaseImpl,
	withdrawUseCase *usecases.WithdrawUseCaseImpl,
	holdUseCase *usecases.HoldUseCaseImpl,
	payUseCase *usecases.PayUseCaseImpl,
) *PaymentHandler {
	return &PaymentHandler{
		depositUseCase:  depositUseCase,
		withdrawUseCase: withdrawUseCase,
		holdUseCase:     holdUseCase,
		payUseCase:      payUseCase,
	}
}

func (h *PaymentHandler) Deposit(ctx context.Context, request *pb.DepositRequest) (*pb.DepositResponse, error) {
	if request.Amount == nil {
		return &pb.DepositResponse{Success: false}, errors.New("error")
	}
	amountStr := request.Amount.GetValue()

	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		log.Println("error: invalid amount format:", err)
		return &pb.DepositResponse{Success: false}, err
	}

	payment := models.Payment{
		UserId:    request.UserId,
		Amount:    amount,
		Type:      "deposit",
		CreatedAt: time.Now(),
	}

	err = h.depositUseCase.Invoke(ctx, payment)
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

	payment := models.Payment{
		UserId:    request.UserId,
		Amount:    amount,
		Type:      "withdraw",
		CreatedAt: time.Now(),
	}

	err = h.withdrawUseCase.Invoke(ctx, payment)
	if err != nil {
		log.Println("withdraw error:", err)
		return &pb.WithDrawResponse{Success: false}, err
	}

	return &pb.WithDrawResponse{Success: true}, nil
}

func (h *PaymentHandler) Hold(ctx context.Context, request *pb.HoldRequest) (*pb.HoldResponse, error) {
	if request.Amount == nil {
		return &pb.HoldResponse{Success: false}, errors.New("amount is required")
	}

	amountStr := request.Amount.GetValue()
	amount, err := decimal.NewFromString(amountStr)
	if err != nil {
		log.Println("Invalid amount format:", err)
		return &pb.HoldResponse{Success: false}, err
	}

	payment := models.Payment{
		UserId:    request.UserId,
		Amount:    amount,
		Type:      "hold",
		CreatedAt: time.Now(),
	}

	err = h.holdUseCase.Invoke(ctx, payment)
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

	err = h.payUseCase.Invoke(ctx, int64(request.UserId), int64(request.LandlordId), rentAmount, pledgeAmount)
	if err != nil {
		log.Println("pay error:", err)
		return &pb.PayResponse{Success: false}, err
	}

	return &pb.PayResponse{Success: true}, nil
}
