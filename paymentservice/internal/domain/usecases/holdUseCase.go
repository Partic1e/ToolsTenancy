package usecases

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	"paymentservice/internal/data/repository"
)

type HoldUseCaseImpl struct {
	repo repository.PaymentRepositoryImpl
}

func NewHoldUseCase(repo repository.PaymentRepositoryImpl) *HoldUseCaseImpl {
	return &HoldUseCaseImpl{repo: repo}
}

func (u *HoldUseCaseImpl) Invoke(ctx context.Context, rentId int64, amount decimal.Decimal) error {
	balance, err := u.repo.GetBalance(ctx, rentId)
	if err != nil {
		return err
	}

	if balance.LessThan(amount) {
		return errors.New("insufficient funds for hold")
	}

	return u.repo.HoldAmount(ctx, rentId, amount)
}
