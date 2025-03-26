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
	createUseCase           *usecase.CreateAdUseCase
	updateUseCase           *usecase.UpdateAdUseCase
	deleteUseCase           *usecase.DeleteAdUseCase
	getCategoriesUseCase    *usecase.GetCategoriesUseCase
	getAdsByCategoryUseCase *usecase.GetAdsByCategoryUseCase
}

func NewAdHandler(
	createUseCase *usecase.CreateAdUseCase,
	updateUseCase *usecase.UpdateAdUseCase,
	deleteUseCase *usecase.DeleteAdUseCase,
	getCategoriesUseCase *usecase.GetCategoriesUseCase,
	getAdsByCategoryUseCase *usecase.GetAdsByCategoryUseCase,
) *AdHandler {
	return &AdHandler{
		createUseCase:           createUseCase,
		updateUseCase:           updateUseCase,
		deleteUseCase:           deleteUseCase,
		getCategoriesUseCase:    getCategoriesUseCase,
		getAdsByCategoryUseCase: getAdsByCategoryUseCase,
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

func (h *AdHandler) UpdateAd(ctx context.Context, req *pb.UpdateAdRequest) (*pb.UpdateAdResponse, error) {
	_, err := decimal.NewFromString(req.CostPerDay)
	if err != nil {
		return nil, fmt.Errorf("неверный формат стоимости в день: %v", err)
	}

	_, err = decimal.NewFromString(req.Deposit)
	if err != nil {
		return nil, fmt.Errorf("неверный формат депозита: %v", err)
	}

	// Вызываем метод UpdateAdUseCase для обновления объявления
	err = h.updateUseCase.UpdateAd(req.Name, req.Description, req.CostPerDay, req.Deposit, req.PhotoPath, req.Id, req.LandlordId, req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка при обновлении объявления: %v", err)
	}

	// Возвращаем обновленные данные
	return &pb.UpdateAdResponse{
		Id:          req.Id,
		Name:        req.Name,
		Description: req.Description,
		CostPerDay:  req.CostPerDay,
		Deposit:     req.Deposit,
		PhotoPath:   req.PhotoPath,
		LandlordId:  req.LandlordId,
		CategoryId:  req.CategoryId,
	}, nil
}

func (h *AdHandler) GetAllCategories(ctx context.Context, req *pb.Empty) (*pb.CategoryList, error) {
	categories, err := h.getCategoriesUseCase.GetAllCategories()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка при получении категорий: %v", err)
	}

	var categoryList []*pb.Category
	for _, c := range categories {
		categoryList = append(categoryList, &pb.Category{
			Id:   c.ID,
			Name: c.Name,
		})
	}

	return &pb.CategoryList{Categories: categoryList}, nil
}

func (h *AdHandler) GetAdsByCategory(ctx context.Context, req *pb.GetAdsByCategoryRequest) (*pb.GetAdsByCategoryResponse, error) {
	ads, err := h.getAdsByCategoryUseCase.GetAdsByCategory(req.CategoryId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "ошибка при получении объявлений: %v", err)
	}

	var pbAds []*pb.Ad
	for _, ad := range ads {
		pbAds = append(pbAds, &pb.Ad{
			Id:          ad.ID,
			Name:        ad.Name,
			Description: ad.Description,
			CostPerDay:  ad.CostPerDay.String(),
			Deposit:     ad.Deposit.String(),
			PhotoPath:   ad.PhotoPath,
			LandlordId:  ad.LandlordId,
			CategoryId:  ad.CategoryId,
		})
	}

	return &pb.GetAdsByCategoryResponse{Ads: pbAds}, nil
}
