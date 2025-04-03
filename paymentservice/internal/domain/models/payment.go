package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Payment struct {
	Id        int64           `json:"id"`
	UserId    int64           `json:"tg_id"`
	Amount    decimal.Decimal `json:"amount"`
	Type      string          `json:"type"`
	CreatedAt time.Time       `json:"created_at"`
}
