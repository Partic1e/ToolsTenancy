package usecase

import (
	"rentservice/internal/repository"
	"time"
)

type GetRentedDatesUseCase struct {
	rentRepo *repository.RentRepository
}

func NewGetRentedDatesUseCase(rentRepo *repository.RentRepository) *GetRentedDatesUseCase {
	return &GetRentedDatesUseCase{rentRepo: rentRepo}
}

func (uc *GetRentedDatesUseCase) GetRentedDates(adID int64) ([]time.Time, error) {
	return uc.rentRepo.GetRentedDates(adID)
}
