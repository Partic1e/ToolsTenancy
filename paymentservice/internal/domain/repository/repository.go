package repository

import (
	"context"
	"github.com/shopspring/decimal"
	"paymentservice/internal/domain/models"
)

type PaymentRepository interface {
	Deposit(ctx context.Context, userID int64, amount decimal.Decimal) error
	Withdraw(ctx context.Context, userID int64, amount decimal.Decimal) error
	HoldFunds(ctx context.Context, userID int64, rentAmount, pledgeAmount decimal.Decimal) error
	ReleaseHeldFunds(ctx context.Context, renterID, heldFundsID, landlordID int64, rentAmount, pledgeAmount decimal.Decimal, toLandlord bool) error
	GetUserBalance(ctx context.Context, userID int64) (decimal.Decimal, error)
	CreateTransaction(ctx context.Context, payment models.Payment) error
}
