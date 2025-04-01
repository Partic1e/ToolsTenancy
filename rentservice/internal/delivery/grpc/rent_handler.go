package grpc

import (
	"context"
	"fmt"
	pb "rentservice/api/rent"
	"rentservice/internal/core/usecase"
)

type RentHandler struct {
	pb.UnimplementedRentServiceServer
	getRentsByLandlordUseCase *usecase.GetRentsByLandlordUseCase
	getRentsByRenterUseCase   *usecase.GetRentsByRenterUseCase
}

func NewRentHandler(
	getRentsByLandlordUseCase *usecase.GetRentsByLandlordUseCase,
	getRentsByRenterUseCase *usecase.GetRentsByRenterUseCase,
) *RentHandler {
	return &RentHandler{
		getRentsByLandlordUseCase: getRentsByLandlordUseCase,
		getRentsByRenterUseCase:   getRentsByRenterUseCase,
	}
}

func (h *RentHandler) GetRentsByLandlord(ctx context.Context, req *pb.GetRentByLandlordRequest) (*pb.GetResponse, error) {
	rents, err := h.getRentsByLandlordUseCase.GetRentsByLandlord(req.LandlordId)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении аренд: %v", err)
	}

	var pbRents []*pb.Rent
	for _, rent := range rents {
		pbRents = append(pbRents, &pb.Rent{
			Id:         rent.ID,
			Status:     rent.Status,
			Cost:       rent.Cost.String(),
			DateStart:  rent.DateStart,
			DateEnd:    rent.DateEnd,
			AdId:       rent.AdID,
			LandlordId: rent.LandlordID,
			RenterId:   rent.RenterID,
		})
	}

	return &pb.GetResponse{Rents: pbRents}, nil
}

func (h *RentHandler) GetRentsByRenter(ctx context.Context, req *pb.GetRentByRenterRequest) (*pb.GetResponse, error) {
	rents, err := h.getRentsByRenterUseCase.GetRentsByRenter(req.RenterId)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении аренд: %v", err)
	}

	var rentList []*pb.Rent
	for _, rent := range rents {
		rentList = append(rentList, &pb.Rent{
			Id:         rent.ID,
			Status:     rent.Status,
			Cost:       rent.Cost.String(),
			DateStart:  rent.DateStart,
			DateEnd:    rent.DateEnd,
			AdId:       rent.AdID,
			LandlordId: rent.LandlordID,
			RenterId:   rent.RenterID,
		})
	}

	return &pb.GetResponse{Rents: rentList}, nil
}
