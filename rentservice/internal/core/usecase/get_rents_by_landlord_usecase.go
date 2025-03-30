package usecase

import (
	"rentservice/internal/core/entity"
	"rentservice/internal/repository"
)

type RentUseCase struct {
	rentRepo *repository.RentRepository
}

func NewRentUseCase(rentRepo *repository.RentRepository) *RentUseCase {
	return &RentUseCase{rentRepo: rentRepo}
}

func (uc *RentUseCase) GetRentsByLandlord(landlordID int64) ([]*entity.Rent, error) {
	return uc.rentRepo.GetRentsByLandlord(landlordID)
}
