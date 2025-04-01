package usecase

import (
	"rentservice/internal/core/entity"
	"rentservice/internal/repository"
)

type GetRentsByLandlordUseCase struct {
	rentRepo *repository.RentRepository
}

func NewGetRentsByLandlordUseCase(rentRepo *repository.RentRepository) *GetRentsByLandlordUseCase {
	return &GetRentsByLandlordUseCase{rentRepo: rentRepo}
}

func (uc *GetRentsByLandlordUseCase) GetRentsByLandlord(landlordID int64) ([]*entity.Rent, error) {
	return uc.rentRepo.GetRentsByLandlord(landlordID)
}
