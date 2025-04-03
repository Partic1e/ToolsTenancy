package usecase

import (
	"context"
	"errors"
	"log"
	"rentservice/api/payment"
	"rentservice/internal/core/entity"
	"rentservice/internal/repository"

	"github.com/shopspring/decimal"
)

type CrateRentUseCase struct {
	rentRepo      *repository.RentRepository
	paymentClient payment.PaymentServiceClient
}

func NewRentUseCase(rentRepo *repository.RentRepository, paymentClient payment.PaymentServiceClient) *CrateRentUseCase {
	return &CrateRentUseCase{rentRepo: rentRepo, paymentClient: paymentClient}
}

func (uc *CrateRentUseCase) CreateRent(ctx context.Context, rentAmount, pledgeAmount, dateStart, dateEnd string, adID, landlordID, renterID int64) (bool, error) {
	rentAmountDec, err := decimal.NewFromString(rentAmount)
	if err != nil {
		return false, errors.New("invalid rent amount format")
	}

	_, err = decimal.NewFromString(pledgeAmount)
	if err != nil {
		return false, errors.New("invalid pledge amount format")
	}

	holdResp, err := uc.paymentClient.Hold(ctx, &payment.HoldRequest{
		RenterId:     renterID,
		RentAmount:   rentAmount,
		PledgeAmount: pledgeAmount,
	})
	if err != nil || !holdResp.Success {
		log.Printf("Hold failed: %v", err)
		return false, nil
	}

	err = uc.rentRepo.CreateRent(entity.Rent{
		Status:     "active",
		Cost:       rentAmountDec,
		DateStart:  dateStart,
		DateEnd:    dateEnd,
		AdID:       adID,
		LandlordID: landlordID,
		RenterID:   renterID,
		HeldID:     holdResp.HeldFundsID,
	})
	if err != nil {
		log.Printf("Failed to create rent: %v", err)
		return false, err
	}

	return true, nil
}
