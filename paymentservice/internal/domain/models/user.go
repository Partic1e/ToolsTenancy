package models

import "github.com/shopspring/decimal"

type User struct {
	TgId    int64           `json:"tg_id"`
	Balance decimal.Decimal `json:"balance"`
	Email   string          `json:"email"`
}
