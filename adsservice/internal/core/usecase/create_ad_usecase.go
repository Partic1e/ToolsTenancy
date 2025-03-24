package usecase

import (
	"adsservice/internal/core/entity"
	"adsservice/internal/repository"
	"github.com/shopspring/decimal"
)

type AdUseCase struct {
	adRepo *repository.AdRepository
}

func NewAdUseCase(adRepo *repository.AdRepository) *AdUseCase {
	return &AdUseCase{adRepo: adRepo}
}

func (uc *AdUseCase) CreateAd(name, description string, costPerDay, deposit decimal.Decimal, photoPath string, landlordId int64) (*entity.Ad, error) {
	ad := &entity.Ad{
		Name:        name,
		Description: description,
		CostPerDay:  costPerDay,
		Deposit:     deposit,
		PhotoPath:   photoPath,
		LandlordId:  landlordId,
	}

	return uc.adRepo.CreateAd(ad)
}
