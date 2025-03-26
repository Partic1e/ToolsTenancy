package usecase

import (
	"adsservice/internal/core/entity"
	"adsservice/internal/repository"
	"github.com/shopspring/decimal"
)

type CreateAdUseCase struct {
	adRepo *repository.AdRepository
}

func NewCreateAdUseCase(adRepo *repository.AdRepository) *CreateAdUseCase {
	return &CreateAdUseCase{adRepo: adRepo}
}

func (uc *CreateAdUseCase) CreateAd(name, description string, costPerDay, deposit decimal.Decimal, photoPath string, landlordId, categoryId int64) (*entity.Ad, error) {
	ad := &entity.Ad{
		Name:        name,
		Description: description,
		CostPerDay:  costPerDay,
		Deposit:     deposit,
		PhotoPath:   photoPath,
		LandlordId:  landlordId,
		CategoryId:  categoryId,
	}

	return uc.adRepo.CreateAd(ad)
}
