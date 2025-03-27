package entity

import "github.com/shopspring/decimal"

type User struct {
	ID      int64
	TgID    int64
	Balance decimal.Decimal
	Email   *string
}
