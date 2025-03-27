package usecase

import "adsservice/internal/repository"

type DeleteAdUseCase struct {
	adRepo *repository.AdRepository
}

func NewDeleteAdUseCase(adRepo *repository.AdRepository) *DeleteAdUseCase {
	return &DeleteAdUseCase{adRepo: adRepo}
}

func (uc *DeleteAdUseCase) DeleteAd(name string, landlordId int64) error {
	return uc.adRepo.DeleteAd(name, landlordId)
}
