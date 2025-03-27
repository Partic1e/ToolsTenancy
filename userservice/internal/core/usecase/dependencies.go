package usecase

import (
	"userservice/internal/repository"
)

type Dependencies struct {
	UserUsecase *UserUseCase
}

func NewDependencies(repo *repository.UserRepository) *Dependencies {
	return &Dependencies{
		UserUsecase: NewUserUseCase(repo),
	}
}
