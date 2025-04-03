package usecases

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"log"
	"paymentservice/internal/data/repository"
	"paymentservice/internal/domain/models"
	"time"
)

type PaymentUseCase interface {
	Deposit(ctx context.Context, userID int64, amount decimal.Decimal) error
	Withdraw(ctx context.Context, userID int64, amount decimal.Decimal) error
	HoldFunds(ctx context.Context, userID int64, amount decimal.Decimal, rentID int64) error
	CompleteRent(ctx context.Context, rentID int64, toLandlord bool) error
	PayRent(ctx context.Context, userID, landlordID int64, rentAmount, pledgeAmount decimal.Decimal) error
}

type PaymentUseCaseImpl struct {
	repo repository.PaymentRepositoryImpl
}

func NewPaymentUseCase(repo repository.PaymentRepositoryImpl) *PaymentUseCaseImpl {
	return &PaymentUseCaseImpl{repo: repo}
}

func (uc *PaymentUseCaseImpl) Deposit(
	ctx context.Context,
	userID int64,
	amount decimal.Decimal,
) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount must be greater than zero")
	}

	err := uc.repo.Deposit(ctx, userID, amount)
	if err != nil {
		return err
	}

	log.Printf("User %d deposited: %s", userID, amount)
	return uc.repo.CreateTransaction(ctx, models.Payment{
		UserId:    userID,
		Amount:    amount,
		Type:      "deposit",
		CreatedAt: time.Now(),
	})
}

func (uc *PaymentUseCaseImpl) Withdraw(
	ctx context.Context,
	userID int64,
	amount decimal.Decimal,
) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount must be greater than zero")
	}

	err := uc.repo.Withdraw(ctx, userID, amount)
	if err != nil {
		return err
	}

	log.Printf("User %d withdrew: %s", userID, amount)
	return uc.repo.CreateTransaction(ctx, models.Payment{
		UserId:    userID,
		Amount:    amount,
		Type:      "withdraw",
		CreatedAt: time.Now(),
	})
}

func (uc *PaymentUseCaseImpl) HoldFunds(
	ctx context.Context,
	userID int64,
	rentAmount,
	pledgeAmount decimal.Decimal,
) (int64, error) {
	if rentAmount.LessThanOrEqual(decimal.Zero) || pledgeAmount.LessThanOrEqual(decimal.Zero) {
		return -1, errors.New("rent and pledge amounts must be greater than zero")
	}

	heldFundsID, err := uc.repo.HoldFunds(ctx, userID, rentAmount, pledgeAmount)
	if err != nil {
		return -1, err
	}

	log.Printf("Held funds: rent %s, pledge %s for user %d", rentAmount, pledgeAmount, userID)

	err = uc.repo.CreateTransaction(ctx, models.Payment{
		UserId:    userID,
		Amount:    rentAmount.Add(pledgeAmount),
		Type:      "hold",
		CreatedAt: time.Now(),
	})
	if err != nil {
		return -1, err
	}

	return heldFundsID, nil
}

func (uc *PaymentUseCaseImpl) CompleteRent(
	ctx context.Context,
	renterID,
	heldFundsID,
	landlordID int64,
	rentAmount,
	pledgeAmount decimal.Decimal,
	toLandlord bool,
) error {
	err := uc.repo.ReleaseHeldFunds(ctx, renterID, heldFundsID, landlordID, rentAmount, pledgeAmount, toLandlord)
	if err != nil {
		return err
	}

	log.Printf("Rent completed. Rent %s, pledge %s transferred. To landlord: %t", rentAmount, pledgeAmount, toLandlord)

	return uc.repo.CreateTransaction(ctx, models.Payment{
		UserId:    renterID,
		Amount:    rentAmount.Add(pledgeAmount),
		Type:      "release",
		CreatedAt: time.Now(),
	})
}
