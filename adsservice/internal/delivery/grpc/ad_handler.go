package grpc

import (
	pb "adsservice/api/ad"
	"adsservice/internal/core/usecase"
	"context"
	"fmt"
	"github.com/shopspring/decimal"
)

type AdHandler struct {
	pb.UnimplementedAdServiceServer
	createUseCase *usecase.CreateAdUseCase
}

func NewAdHandler(
	createUseCase *usecase.CreateAdUseCase,
) *AdHandler {
	return &AdHandler{
		createUseCase: createUseCase,
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
