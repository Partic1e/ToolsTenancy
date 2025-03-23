package usecases

import (
	"errors"
	"paymentservice/internal/domain/models"
	"paymentservice/internal/repository"
)

type WithdrawUseCase interface {
	Invoke(payment models.Payment) error
}

type WithdrawUseCaseImpl struct {
	repository repository.PaymentRepositoryImpl
}

func NewWithdrawUseCase(repository repository.PaymentRepositoryImpl) *WithdrawUseCaseImpl {
	return &WithdrawUseCaseImpl{repository: repository}
}

func (w *WithdrawUseCaseImpl) Invoke(payment models.Payment) error {
	user, err := w.repository.GetUserById(payment.UserId)
	if err != nil {
		return err
	}

	if user.Balance.LessThan(payment.Amount) {
		return errors.New("insufficient funds")
	}

	user.Balance = user.Balance.Sub(payment.Amount)
	err = w.repository.UpdateBalance(user.Id, user.Balance)
	if err != nil {
		return err
	}

	return w.repository.CreatePayment(payment)
}
