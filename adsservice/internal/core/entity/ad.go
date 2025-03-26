package entity

import "github.com/shopspring/decimal"

type Ad struct {
	ID          int64
	Name        string
	Description string
	CostPerDay  decimal.Decimal
	Deposit     decimal.Decimal
	PhotoPath   string
	LandlordId  int64
	CategoryId  int64
}
