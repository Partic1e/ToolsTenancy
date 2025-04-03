package repository

import (
	"context"
	"github.com/shopspring/decimal"
	"paymentservice/internal/domain/models"
)

type PaymentRepository interface {
	Deposit(ctx context.Context, userID int64, amount decimal.Decimal) error
	Withdraw(ctx context.Context, userID int64, amount decimal.Decimal) error
	HoldFunds(ctx context.Context, userID int64, amount decimal.Decimal, rentID int64) error
	ReleaseHeldFunds(ctx context.Context, rentID int64, toLandlord bool) error
	TransferFunds(ctx context.Context, fromUserID, toUserID int64, amount decimal.Decimal) error
	GetUserBalance(ctx context.Context, userID int64) (decimal.Decimal, error)
	GetHeldAmount(ctx context.Context, rentID int64) (decimal.Decimal, error)
	CreateTransaction(ctx context.Context, payment models.Payment) error
}
