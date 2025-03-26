package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/shopspring/decimal"
	"userservice/internal/core/entity"
)

type UserRepositoryInterface interface {
	GetOrCreateUser(tgID int64) (*entity.User, error)
	//Реализовать, если/когда будет сервис уведомлений - Привязка почты.
	//AddOrUpdateUserEmail(email string) (*entity.User, error)
}

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetOrCreateUser(tgID int64) (*entity.User, error) {
	var user entity.User
	var balanceStr string

	err := r.db.QueryRow("SELECT tg_id, balance, email FROM users WHERE tg_id = $1", tgID).
		Scan(&user.TgID, &balanceStr, &user.Email)

	if errors.Is(err, sql.ErrNoRows) {
		defaultBalance := decimal.NewFromFloat(0.00)
		err = r.db.QueryRow(
			"INSERT INTO users (tg_id, balance) VALUES ($1, $2) RETURNING tg_id, balance, email",
			tgID, defaultBalance).
			Scan(&user.TgID, &balanceStr, &user.Email)

		if err != nil {
			return nil, fmt.Errorf("[UserService][Postgres] ошибка при создании пользователя: %v", err)
		}

		log.Printf("[UserService][Postgres] Создан новый пользователь с tg_id: %d", tgID)
	} else if err != nil {
		return nil, fmt.Errorf("[UserService][Postgres] ошибка при поиске пользователя: %v", err)
	}

	user.Balance, err = decimal.NewFromString(balanceStr)
	if err != nil {
		return nil, fmt.Errorf("[UserService][Postgres] ошибка преобразования баланса: %v", err)
	}

	return &user, nil
}
