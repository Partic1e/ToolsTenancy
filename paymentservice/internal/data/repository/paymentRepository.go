package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/shopspring/decimal"
	"paymentservice/internal/domain/models"
)

type PaymentRepositoryImpl struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{db: db}
}

func (r *PaymentRepositoryImpl) Deposit(ctx context.Context, userID int64, amount decimal.Decimal) error {
	_, err := r.db.ExecContext(ctx, "UPDATE Users SET balance = balance + ? WHERE id = ?", amount.String(), userID)
	return err
}

func (r *PaymentRepositoryImpl) Withdraw(ctx context.Context, userID int64, amount decimal.Decimal) error {
	balance, err := r.GetUserBalance(ctx, userID)
	if err != nil {
		return err
	}

	if balance.LessThan(amount) {
		return errors.New("insufficient funds")
	}

	_, err = r.db.ExecContext(ctx, "UPDATE Users SET balance = balance - ? WHERE id = ?", amount.String(), userID)
	return err
}

func (r *PaymentRepositoryImpl) HoldFunds(ctx context.Context, userID int64, amount decimal.Decimal, rentID int64) error {
	balance, err := r.GetUserBalance(ctx, userID)
	if err != nil {
		return err
	}

	if balance.LessThan(amount) {
		return errors.New("insufficient funds")
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance - ? WHERE id = ?", amount.String(), userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO WithheldFunds (amount, rent_id) VALUES (?, ?)", amount.String(), rentID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PaymentRepositoryImpl) ReleaseHeldFunds(ctx context.Context, rentID int64, toLandlord bool) error {
	var amount decimal.Decimal
	var renterID, landlordID int64

	err := r.db.QueryRowContext(ctx, `
		SELECT wf.amount, r.renter_id, r.landlord_id
		FROM WithheldFunds wf
		JOIN Rents r ON wf.rent_id = r.id
		WHERE wf.rent_id = ?`, rentID).Scan(&amount, &renterID, &landlordID)
	if err != nil {
		return err
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if toLandlord {
		_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance + ? WHERE id = ?", amount.String(), landlordID)
	} else {
		_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance + ? WHERE id = ?", amount.String(), renterID)
	}
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM WithheldFunds WHERE rent_id = ?", rentID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PaymentRepositoryImpl) TransferFunds(ctx context.Context, fromUserID, toUserID int64, amount decimal.Decimal) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance - ? WHERE id = ?", amount.String(), fromUserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance + ? WHERE id = ?", amount.String(), toUserID)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *PaymentRepositoryImpl) GetUserBalance(ctx context.Context, userID int64) (decimal.Decimal, error) {
	var balance decimal.Decimal
	err := r.db.QueryRowContext(ctx, "SELECT balance FROM Users WHERE id = ?", userID).Scan(&balance)
	if err != nil {
		return decimal.Zero, err
	}
	return balance, nil
}

func (r *PaymentRepositoryImpl) GetHeldAmount(ctx context.Context, rentID int64) (decimal.Decimal, error) {
	var amount decimal.Decimal
	err := r.db.QueryRowContext(ctx, "SELECT amount FROM WithheldFunds WHERE rent_id = ?", rentID).Scan(&amount)
	if err != nil {
		return decimal.Zero, err
	}
	return amount, nil
}

func (r *PaymentRepositoryImpl) CreateTransaction(ctx context.Context, payment models.Payment) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO Payments (user_id, amount, type, created_at) 
		VALUES (?, ?, ?, ?)`, payment.UserId, payment.Amount.String(), payment.Type, payment.CreatedAt)
	return err
}
