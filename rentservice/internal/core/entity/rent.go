package entity

import "github.com/shopspring/decimal"

type Rent struct {
	ID         int64           `json:"id"`
	Status     string          `json:"status"`
	Cost       decimal.Decimal `json:"cost"`
	DateStart  string          `json:"date_start"`
	DateEnd    string          `json:"date_end"`
	AdID       int64           `json:"ad_id"`
	LandlordID int64           `json:"landlord_id"`
	RenterID   int64           `json:"renter_id"`
	HeldID     int64           `json:"held_id"`
}
