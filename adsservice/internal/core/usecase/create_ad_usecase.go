package usecase

import (
	"adsservice/internal/core/entity"
	"adsservice/internal/repository"
)

type CreateAdUseCase struct {
	adRepo *repository.AdRepository
}

func NewCreateAdUseCase(adRepo *repository.AdRepository) *CreateAdUseCase {
	return &CreateAdUseCase{adRepo: adRepo}
}

func (uc *CreateAdUseCase) CreateAd(ad *entity.Ad) (*entity.Ad, error) {
	return uc.adRepo.CreateAd(ad)
}
