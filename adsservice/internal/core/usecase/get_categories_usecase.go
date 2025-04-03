package usecase

import (
	"adsservice/internal/core/entity"
	"adsservice/internal/repository"
)

type GetCategoriesUseCase struct {
	adRepo *repository.AdRepository
}

func NewGetCategoriesUseCase(adRepo *repository.AdRepository) *GetCategoriesUseCase {
	return &GetCategoriesUseCase{adRepo: adRepo}
}

func (uc *GetCategoriesUseCase) GetAllCategories() ([]entity.Category, error) {
	return uc.adRepo.GetAllCategories()
}
