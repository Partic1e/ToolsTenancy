package models

import (
	"github.com/shopspring/decimal"
)

type Payment struct {
	Id        uint64          `json:"id"`
	UserId    uint64          `json:"user_id"`
	Amount    decimal.Decimal `json:"amount"`
	Type      string          `json:"type"`
	CreatedAt int64           `json:"created_at"`
}
