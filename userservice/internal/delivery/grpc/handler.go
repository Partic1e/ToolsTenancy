package grpc

import (
	"context"
	"log"

	_ "github.com/shopspring/decimal"
	pb "userservice/api/user"
	"userservice/internal/core/usecase"
)

type UserHandler struct {
	pb.UnimplementedUserServiceServer
	userUseCase *usecase.UserUseCase
}

func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase: userUseCase}
}

func (h *UserHandler) GetOrCreateUser(ctx context.Context, req *pb.GetOrCreateUserRequest) (*pb.GetOrCreateUserResponse, error) {
	user, err := h.userUseCase.GetOrCreateUser(req.TgId)
	if err != nil {
		log.Printf("Ошибка получения пользователя: %v", err)
		return nil, err
	}
	return &pb.GetOrCreateUserResponse{
		TgId:    user.TgID,
		Balance: user.Balance.String(),
	}, nil
}
