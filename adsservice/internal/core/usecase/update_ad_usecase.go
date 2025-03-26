package usecase

import (
	"adsservice/internal/repository"
	"fmt"
	"github.com/shopspring/decimal"
)

type UpdateAdUseCase struct {
	adRepo *repository.AdRepository
}

func NewUpdateAdUseCase(adRepo *repository.AdRepository) *UpdateAdUseCase {
	return &UpdateAdUseCase{adRepo: adRepo}
}

func (u *UpdateAdUseCase) UpdateAd(name, description, costPerDay, deposit, photoPath string, id, landlordId, categoryId int64) error {
	costPerDayDecimal, err := decimal.NewFromString(costPerDay)
	if err != nil {
		return fmt.Errorf("неверный формат стоимости в день: %v", err)
	}

	depositDecimal, err := decimal.NewFromString(deposit)
	if err != nil {
		return fmt.Errorf("неверный формат депозита: %v", err)
	}

	err = u.adRepo.UpdateAd(name, description, costPerDayDecimal, depositDecimal, photoPath, id, landlordId, categoryId)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении объявления: %v", err)
	}

	return nil
}
