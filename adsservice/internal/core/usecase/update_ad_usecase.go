package usecase

import (
	"adsservice/internal/core/entity"
	"adsservice/internal/repository"
)

type UpdateAdUseCase struct {
	adRepo *repository.AdRepository
}

func NewUpdateAdUseCase(adRepo *repository.AdRepository) *UpdateAdUseCase {
	return &UpdateAdUseCase{adRepo: adRepo}
}

func (u *UpdateAdUseCase) UpdateAd(ad *entity.Ad) error {
	return u.adRepo.UpdateAd(ad)
}
