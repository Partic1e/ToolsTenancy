package usecases

import (
	"paymentservice/internal/domain/models"
	"paymentservice/internal/repository"
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

func (d *DepositUseCaseImpl) Invoke(payment models.Payment) error {
	user, err := d.repository.GetUserById(payment.UserId)
	if err != nil {
		return err
	}

	user.Balance = user.Balance.Add(payment.Amount)
	err = d.repository.UpdateBalance(user.Id, user.Balance)
	if err != nil {
		return err
	}

	return d.repository.CreatePayment(payment)
}
