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
	adUseCase *usecase.AdUseCase
}

func NewAdHandler(adUseCase *usecase.AdUseCase) *AdHandler {
	return &AdHandler{adUseCase: adUseCase}
}

func (h *AdHandler) CreateAd(ctx context.Context, req *pb.CreateAdRequest) (*pb.AdResponse, error) {
	costPerDay, err := decimal.NewFromString(req.CostPerDay)
	if err != nil {
		return nil, fmt.Errorf("неверный формат стоимости в день")
	}

	deposit, err := decimal.NewFromString(req.Deposit)
	if err != nil {
		return nil, fmt.Errorf("неверный формат депозита")
	}

	ad, err := h.adUseCase.CreateAd(
		req.Name, req.Description, costPerDay, deposit, req.PhotoPath, req.LandlordId,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании объявления: %v", err)
	}

	return &pb.AdResponse{
		Id:          ad.ID,
		Name:        ad.Name,
		Description: ad.Description,
		CostPerDay:  ad.CostPerDay.String(),
		Deposit:     ad.Deposit.String(),
		PhotoPath:   ad.PhotoPath,
		LandlordId:  ad.LandlordId,
	}, nil
}
