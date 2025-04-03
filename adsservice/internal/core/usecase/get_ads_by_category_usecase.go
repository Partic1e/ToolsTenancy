package usecase

import (
	"adsservice/internal/core/entity"
	"adsservice/internal/repository"
)

type GetAdsByCategoryUseCase struct {
	adRepo *repository.AdRepository
}

func NewGetAdsByCategoryUseCase(adRepo *repository.AdRepository) *GetAdsByCategoryUseCase {
	return &GetAdsByCategoryUseCase{adRepo: adRepo}
}

func (uc *GetAdsByCategoryUseCase) GetAdsByCategory(categoryID int64) ([]*entity.Ad, error) {
	return uc.adRepo.GetAdsByCategory(categoryID)
}
