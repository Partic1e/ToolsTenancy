package usecase

import (
	"rentservice/internal/core/entity"
	"rentservice/internal/repository"
)

type GetRentsByRenterUseCase struct {
	rentRepo *repository.RentRepository
}

func NewGetRentsByRenterUseCase(rentRepo *repository.RentRepository) *GetRentsByRenterUseCase {
	return &GetRentsByRenterUseCase{rentRepo: rentRepo}
}

func (uc *GetRentsByRenterUseCase) GetRentsByRenter(renterId int64) ([]*entity.Rent, error) {
	return uc.rentRepo.GetRentsByRenter(renterId)
}
