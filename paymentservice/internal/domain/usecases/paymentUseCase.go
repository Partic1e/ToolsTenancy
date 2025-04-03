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

func (uc *PaymentUseCaseImpl) Deposit(ctx context.Context, userID int64, amount decimal.Decimal) error {
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

func (uc *PaymentUseCaseImpl) Withdraw(ctx context.Context, userID int64, amount decimal.Decimal) error {
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

func (uc *PaymentUseCaseImpl) HoldFunds(ctx context.Context, userID int64, amount decimal.Decimal, rentID int64) error {
	if amount.LessThanOrEqual(decimal.Zero) {
		return errors.New("amount must be greater than zero")
	}

	err := uc.repo.HoldFunds(ctx, userID, amount, rentID)
	if err != nil {
		return err
	}

	log.Printf("Held funds: %s for rent %d by user %d", amount, rentID, userID)
	return uc.repo.CreateTransaction(ctx, models.Payment{
		UserId:    userID,
		Amount:    amount,
		Type:      "hold",
		CreatedAt: time.Now(),
	})
}

func (uc *PaymentUseCaseImpl) CompleteRent(ctx context.Context, rentID int64, toLandlord bool) error {
	amount, err := uc.repo.GetHeldAmount(ctx, rentID)
	if err != nil {
		return err
	}

	err = uc.repo.ReleaseHeldFunds(ctx, rentID, toLandlord)
	if err != nil {
		return err
	}

	log.Printf("Rent %d completed. Amount %s transferred. To landlord: %t", rentID, amount, toLandlord)
	return uc.repo.CreateTransaction(ctx, models.Payment{
		UserId:    rentID,
		Amount:    amount,
		Type:      "release",
		CreatedAt: time.Now(),
	})
}

func (uc *PaymentUseCaseImpl) PayRent(ctx context.Context, userID, landlordID, rentID int64, rentAmount, pledgeAmount decimal.Decimal) error {
	if rentAmount.LessThanOrEqual(decimal.Zero) || pledgeAmount.LessThanOrEqual(decimal.Zero) {
		return errors.New("rent and pledge amounts must be greater than zero")
	}

	totalAmount := rentAmount.Add(pledgeAmount)

	err := uc.repo.HoldFunds(ctx, userID, totalAmount, rentID)
	if err != nil {
		return err
	}

	log.Printf("User %d paid rent: %s and pledge: %s to landlord %d", userID, rentAmount, pledgeAmount, landlordID)
	return uc.repo.CreateTransaction(ctx, models.Payment{
		UserId:    userID,
		Amount:    totalAmount,
		Type:      "pay_rent",
		CreatedAt: time.Now(),
	})
}
