package usecase

import (
	"userservice/internal/core/entity"
	"userservice/internal/repository"
)

type UserUseCase struct {
	userRepo *repository.UserRepository
}

func NewUserUseCase(userRepo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

func (uc *UserUseCase) GetOrCreateUser(tgID int64) (*entity.User, error) {
	return uc.userRepo.GetOrCreateUser(tgID)
}
