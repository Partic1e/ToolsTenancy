package usecase

import (
	"context"
	"rentservice/api/payment"
	"rentservice/internal/repository"
)

type CloseRentUseCase struct {
	rentRepo      *repository.RentRepository
	paymentClient payment.PaymentServiceClient
}

func NewCloseRentUseCase(rentRepo *repository.RentRepository, paymentClient payment.PaymentServiceClient) *CloseRentUseCase {
	return &CloseRentUseCase{rentRepo: rentRepo, paymentClient: paymentClient}
}

func (uc *CloseRentUseCase) CloseRent(
	ctx context.Context,
	rentID,
	renterID,
	landlordID,
	heldFundsID int64,
	toLandlord bool) (bool, error) {
	completeResp, err := uc.paymentClient.CompleteRent(ctx, &payment.CompleteRentRequest{
		RenterId:    renterID,
		LandlordId:  landlordID,
		HeldFundsID: heldFundsID,
		ToLandlord:  toLandlord,
	})
	if err != nil || !completeResp.Success {
		return false, err
	}

	err = uc.rentRepo.CompleteRent(rentID)
	if err != nil {
		return false, err
	}

	return true, nil
}
