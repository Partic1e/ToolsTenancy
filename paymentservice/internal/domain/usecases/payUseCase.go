package usecases

import (
	"context"
	"errors"
	"paymentservice/internal/data/repository"
)

type PayUseCaseImpl struct {
	repo repository.PaymentRepositoryImpl
}

func NewPayUseCase(repo repository.PaymentRepositoryImpl) *PayUseCaseImpl {
	return &PayUseCaseImpl{repo: repo}
}

func (u *PayUseCaseImpl) Invoke(ctx context.Context, rentId int64) error {
	rent, err := u.repo.GetRentById(ctx, rentId)
	if err != nil {
		return err
	}

	heldBalance, err := u.repo.GetHeldBalance(ctx, rent.RenterId)
	if err != nil {
		return err
	}

	totalAmount := rent.Cost.Add(rent.Deposit)

	if heldBalance.LessThan(totalAmount) {
		return errors.New("insufficient held funds for payment")
	}

	tx, err := u.repo.BeginTransaction(ctx)
	if err != nil {
		return err
	}

	// Перевод аренды
	err = u.repo.Transfer(ctx, rent.RenterId, rent.LandlordId, rent.Cost)
	if err != nil {
		tx.Rollback()
		return err
	}

	// Разблокировка залога
	err = u.repo.ReleaseHold(ctx, rent.RenterId, rent.Deposit)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
