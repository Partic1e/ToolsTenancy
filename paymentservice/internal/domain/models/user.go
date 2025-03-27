package models

import "github.com/shopspring/decimal"

type User struct {
	Id      uint64          `json:"id"`
	TgId    uint64          `json:"tg_id"`
	Balance decimal.Decimal `json:"balance"`
	Email   string          `json:"email"`
}
