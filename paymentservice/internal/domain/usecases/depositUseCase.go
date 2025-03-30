package usecases

import (
	"context"
	"paymentservice/internal/data/repository"
	"paymentservice/internal/domain/models"
)

type DepositUseCase interface {
	Invoke(payment models.Payment) error
}

type DepositUseCaseImpl struct {
	repository repository.PaymentRepositoryImpl
}

func NewDepositUseCase(repository repository.PaymentRepositoryImpl) *DepositUseCaseImpl {
	return &DepositUseCaseImpl{repository: repository}
}

func (d *DepositUseCaseImpl) Invoke(ctx context.Context, payment models.Payment) error {
	user, err := d.repository.GetUserById(ctx, payment.UserId)
	if err != nil {
		return err
	}

	user.Balance = user.Balance.Add(payment.Amount)
	err = d.repository.UpdateBalance(ctx, user.TgId, user.Balance)
	if err != nil {
		return err
	}

	return d.repository.CreatePayment(ctx, payment)
}
