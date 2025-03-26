package usecase

import (
	"adsservice/internal/core/entity"
	"adsservice/internal/repository"
)

type GetAdsByLandlordUseCase struct {
	adRepo *repository.AdRepository
}

func NewGetAdsByLandlordUseCase(adRepo *repository.AdRepository) *GetAdsByLandlordUseCase {
	return &GetAdsByLandlordUseCase{adRepo: adRepo}
}

func (uc *GetAdsByLandlordUseCase) GetAdsByLandlord(landlordID int64) ([]*entity.Ad, error) {
	return uc.adRepo.GetAdsByLandlord(landlordID)
}
