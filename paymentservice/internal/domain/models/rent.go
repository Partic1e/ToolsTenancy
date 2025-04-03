package models

import (
	"github.com/shopspring/decimal"
	"time"
)

type Rent struct {
	Id         int64           `json:"id"`
	Status     string          `json:"status"`
	Cost       decimal.Decimal `json:"cost"`
	DateStart  time.Time       `json:"date_start"`
	DateEnd    time.Time       `json:"date_end"`
	AdId       int64           `json:"ad_id"`
	LandlordId int64           `json:"landlord_id"`
	RenterId   int64           `json:"renter_id"`
	Deposit    decimal.Decimal `json:"deposit"`
}
