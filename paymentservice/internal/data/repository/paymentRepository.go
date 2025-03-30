package repository

import (
	"context"
	"database/sql"
	"github.com/shopspring/decimal"
	"paymentservice/internal/domain/models"
)

type PaymentRepository interface {
	GetUserById(ctx context.Context, userId int64) (*models.User, error)
	UpdateBalance(ctx context.Context, userId int64, newBalance decimal.Decimal) error
	CreatePayment(ctx context.Context, payment models.Payment) error
	GetBalance(ctx context.Context, userId int64) (decimal.Decimal, error)
	HoldAmount(ctx context.Context, rentId int64, amount decimal.Decimal) error
	GetHeldBalance(ctx context.Context, userId int64) (decimal.Decimal, error)
	ReleaseHold(ctx context.Context, userId int64, amount decimal.Decimal) error
	Transfer(ctx context.Context, fromUserId, toUserId int64, amount decimal.Decimal) error
	BeginTransaction(ctx context.Context) (*sql.Tx, error)
	CommitTransaction(tx *sql.Tx) error
	RollbackTransaction(tx *sql.Tx) error
	GetRentById(ctx context.Context, rentId int64) (*models.Rent, error)
}

type PaymentRepositoryImpl struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepositoryImpl {

	return &PaymentRepositoryImpl{db: db}
}

func (p *PaymentRepositoryImpl) GetUserById(ctx context.Context, userId int64) (*models.User, error) {
	var user models.User
	query := "SELECT tg_id, balance, email FROM users WHERE tg_id = $1"
	err := p.db.QueryRowContext(ctx, query, userId).Scan(&user.TgId, &user.Balance, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *PaymentRepositoryImpl) UpdateBalance(ctx context.Context, userId int64, newBalance decimal.Decimal) error {
	query := "UPDATE users SET balance = $2 WHERE tg_id = $1"
	_, err := p.db.ExecContext(ctx, query, userId, newBalance)

	return err
}

func (p *PaymentRepositoryImpl) CreatePayment(ctx context.Context, payment models.Payment) error {
	query := "INSERT INTO payments (user_id, amount, type, created_at) VALUES ($1, $2, $3, $4)"
	_, err := p.db.ExecContext(ctx, query, payment.UserId, payment.Amount, payment.Type, payment.CreatedAt)

	return err
}

func (p *PaymentRepositoryImpl) GetBalance(ctx context.Context, userId int64) (decimal.Decimal, error) {
	var balance decimal.Decimal
	query := "SELECT balance FROM users WHERE tg_id = $1"
	err := p.db.QueryRowContext(ctx, query, userId).Scan(&balance)
	if err != nil {
		return decimal.Zero, err
	}

	return balance, nil
}

func (p *PaymentRepositoryImpl) HoldAmount(ctx context.Context, rentId int64, amount decimal.Decimal) error {
	query := "INSERT INTO withheldfunds (rent_id, amount) VALUES ($1, $2)"
	_, err := p.db.ExecContext(ctx, query, rentId, amount)

	return err
}

func (p *PaymentRepositoryImpl) GetHeldBalance(ctx context.Context, userId int64) (decimal.Decimal, error) {
	var heldAmount decimal.Decimal
	query := `SELECT COALESCE(SUM(amount), 0) FROM withheldfunds 
              	JOIN rents ON withheldfunds.rent_id = rents.id WHERE renter_id = $1`
	err := p.db.QueryRowContext(ctx, query, userId).Scan(&heldAmount)
	if err != nil {
		return decimal.Zero, err
	}

	return heldAmount, nil
}

func (p *PaymentRepositoryImpl) ReleaseHold(ctx context.Context, userId int64, amount decimal.Decimal) error {
	query := `DELETE FROM withheldfunds WHERE rent_id IN 
             	(SELECT id FROM rents WHERE renter_id = $1) AND amount = $2 LIMIT 1`
	_, err := p.db.ExecContext(ctx, query, userId, amount)

	return err
}

func (p *PaymentRepositoryImpl) Transfer(ctx context.Context, fromUserId, toUserId int64, amount decimal.Decimal) error {
	tx, err := p.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := "UPDATE users SET balance = balance - $2 WHERE tg_id = $1"
	_, err = tx.ExecContext(ctx, query, fromUserId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = "UPDATE users SET balance = balance + $2 WHERE tg_id = $1"
	_, err = tx.ExecContext(ctx, query, toUserId, amount)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (p *PaymentRepositoryImpl) BeginTransaction(ctx context.Context) (*sql.Tx, error) {

	return p.db.BeginTx(ctx, nil)
}

func (p *PaymentRepositoryImpl) CommitTransaction(tx *sql.Tx) error {

	return tx.Commit()
}

func (p *PaymentRepositoryImpl) RollbackTransaction(tx *sql.Tx) error {

	return tx.Rollback()
}

func (p *PaymentRepositoryImpl) GetRentById(ctx context.Context, rentId int64) (*models.Rent, error) {
	var rent models.Rent
	query := "SELECT id, cost, deposit, landlord_id, renter_id FROM rents WHERE id = $1"

	err := p.db.QueryRowContext(ctx, query, rentId).Scan(&rent.Id, &rent.Cost, &rent.Deposit, &rent.LandlordId, &rent.RenterId)
	if err != nil {
		return nil, err
	}

	return &rent, nil
}
