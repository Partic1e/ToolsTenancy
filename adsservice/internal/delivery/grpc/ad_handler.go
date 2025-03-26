package grpc

import (
	pb "adsservice/api/ad"
	"adsservice/internal/core/usecase"
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AdHandler struct {
	pb.UnimplementedAdServiceServer
	createUseCase *usecase.CreateAdUseCase
	deleteUseCase *usecase.DeleteAdUseCase
}

func NewAdHandler(
	createUseCase *usecase.CreateAdUseCase,
	deleteUseCase *usecase.DeleteAdUseCase,
) *AdHandler {
	return &AdHandler{
		createUseCase: createUseCase,
		deleteUseCase: deleteUseCase,
	}
}

func (h *AdHandler) CreateAd(ctx context.Context, req *pb.CreateAdRequest) (*pb.CreateAdResponse, error) {
	costPerDay, err := decimal.NewFromString(req.CostPerDay)
	if err != nil {
		return nil, fmt.Errorf("неверный формат стоимости в день: %v", err)
	}

	deposit, err := decimal.NewFromString(req.Deposit)
	if err != nil {
		return nil, fmt.Errorf("неверный формат депозита: %v", err)
	}

	ad, err := h.createUseCase.CreateAd(req.Name, req.Description, costPerDay, deposit, req.PhotoPath, req.LandlordId, req.CategoryId)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании объявления: %v", err)
	}

	return &pb.CreateAdResponse{
		Id:          ad.ID,
		Name:        ad.Name,
		Description: ad.Description,
		CostPerDay:  ad.CostPerDay.String(),
		Deposit:     ad.Deposit.String(),
		PhotoPath:   ad.PhotoPath,
		LandlordId:  ad.LandlordId,
		CategoryId:  ad.CategoryId,
	}, nil
}

func (h *AdHandler) DeleteAd(ctx context.Context, req *pb.DeleteAdRequest) (*pb.DeleteAdResponse, error) {
	err := h.deleteUseCase.DeleteAd(req.Name, req.LandlordId) // Удаляем по названию
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка при удалении объявления: %v", err)
	}
	return &pb.DeleteAdResponse{Success: true}, nil
}
