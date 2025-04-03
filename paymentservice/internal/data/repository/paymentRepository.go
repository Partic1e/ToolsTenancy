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

func (r *PaymentRepositoryImpl) Deposit(
	ctx context.Context,
	userID int64,
	amount decimal.Decimal,
) error {
	_, err := r.db.ExecContext(ctx, "UPDATE Users SET balance = balance + ? WHERE id = ?", amount.String(), userID)
	return err
}

func (r *PaymentRepositoryImpl) Withdraw(
	ctx context.Context,
	userID int64,
	amount decimal.Decimal,
) error {
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

func (r *PaymentRepositoryImpl) HoldFunds(
	ctx context.Context,
	userID int64,
	rentAmount,
	pledgeAmount decimal.Decimal,
) (int64, error) { // ✅ Теперь возвращает heldFundsID
	balance, err := r.GetUserBalance(ctx, userID)
	if err != nil {
		return -1, err
	}

	totalAmount := rentAmount.Add(pledgeAmount)
	if balance.LessThan(totalAmount) {
		return -1, errors.New("insufficient funds")
	}

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return -1, err
	}

	_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance - ? WHERE id = ?", totalAmount.String(), userID)
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	result, err := tx.ExecContext(ctx, "INSERT INTO WithheldFunds (pledge, rent_cost) VALUES (?, ?, ?)", pledgeAmount.String(), rentAmount.String())
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	heldFundsID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return -1, err
	}

	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return heldFundsID, nil
}

func (r *PaymentRepositoryImpl) ReleaseHeldFunds(
	ctx context.Context,
	renterID,
	heldFundsID,
	landlordID int64,
	rentAmount,
	pledgeAmount decimal.Decimal,
	toLandlord bool,
) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if toLandlord {
		totalAmount := rentAmount.Add(pledgeAmount)
		_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance + ? WHERE id = ?", totalAmount.String(), landlordID)
	} else {
		_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance + ? WHERE id = ?", rentAmount.String(), landlordID)
		if err == nil {
			_, err = tx.ExecContext(ctx, "UPDATE Users SET balance = balance + ? WHERE id = ?", pledgeAmount.String(), renterID)
		}
	}

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(ctx, "DELETE FROM WithheldFunds WHERE id = ?", heldFundsID)
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

func (r *PaymentRepositoryImpl) CreateTransaction(ctx context.Context, payment models.Payment) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO Payments (user_id, amount, type, created_at) 
		VALUES (?, ?, ?, ?)`, payment.UserId, payment.Amount.String(), payment.Type, payment.CreatedAt)
	return err
}

func (r *PaymentRepositoryImpl) GetHeldFundsAmount(ctx context.Context, heldFundsID int64) (decimal.Decimal, decimal.Decimal, error) {
	var rentAmount, pledgeAmount decimal.Decimal

	row := r.db.QueryRowContext(ctx, `
        SELECT rent_amount, pledge_amount
        FROM WithheldFunds
        WHERE id = ?`, heldFundsID)

	err := row.Scan(&rentAmount, &pledgeAmount)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}

	return rentAmount, pledgeAmount, nil
}
