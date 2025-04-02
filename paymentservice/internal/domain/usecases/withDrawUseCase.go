package usecases

import (
	"errors"
	"paymentservice/internal/data/repository"
	"paymentservice/internal/domain/models"
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
	err = w.repository.UpdateBalance(user.TgId, user.Balance)
	if err != nil {
		return err
	}

	return w.repository.CreatePayment(payment)
}
