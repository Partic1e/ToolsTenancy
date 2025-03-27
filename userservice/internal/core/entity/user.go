package entity

import "github.com/shopspring/decimal"

type User struct {
	TgID    int64
	Balance decimal.Decimal
	Email   *string
}
