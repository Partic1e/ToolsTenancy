package repository

import (
	"database/sql"
	"fmt"
	"github.com/shopspring/decimal"
	"paymentservice/internal/domain/models"
)

type PaymentRepository interface {
	GetUserById(userId int64) (*models.User, error)
	UpdateBalance(userId int64, newBalance decimal.Decimal) error
	CreatePayment(payment models.Payment) error
}

type PaymentRepositoryImpl struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) *PaymentRepositoryImpl {
	return &PaymentRepositoryImpl{db: db}
}

func (p *PaymentRepositoryImpl) GetUserById(userId int64) (*models.User, error) {
	var user models.User
	query := "SELECT * FROM users WHERE tg_id = $1"
	err := p.db.QueryRow(query, userId).Scan(&user.TgId, &user.Balance, &user.Email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (p *PaymentRepositoryImpl) UpdateBalance(userId int64, newBalance decimal.Decimal) error {
	query := "UPDATE users SET balance = $2 WHERE tg_id = $1"
	_, err := p.db.Exec(query, userId, newBalance)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}

func (p *PaymentRepositoryImpl) CreatePayment(payment models.Payment) error {
	query := "INSERT INTO payments (user_id, amount, type, created_at) VALUES ($1, $2, $3, $4)"
	_, err := p.db.Exec(query, payment.UserId, payment.Amount, payment.Type, payment.CreatedAt)
	if err != nil {
		return fmt.Errorf("error: %w", err)
	}
	return nil
}
