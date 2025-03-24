package repository

import (
	"database/sql"
	"fmt"
	"log"
	"userservice/config"

	_ "github.com/lib/pq"
)

func NewDB(cfg *config.DatabaseConfig) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("[UserService][Postgres] ❌  - ошибка подключения к базе данных: %w", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("[UserService][Postgres] ❌  - ошибка пинга базы данных: %w", err)
	}

	log.Println("[UserService][Postgres] ✅  - Подключен к БД")
	return db, nil
}
