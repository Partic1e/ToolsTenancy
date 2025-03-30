package grpc

import (
	"context"
	"fmt"
	pb "rentservice/api/rent"
	"rentservice/internal/core/usecase"
)

type RentHandler struct {
	pb.UnimplementedRentServiceServer
	rentUseCase *usecase.RentUseCase
}

func NewRentHandler(rentUseCase *usecase.RentUseCase) *RentHandler {
	return &RentHandler{rentUseCase: rentUseCase}
}

func (h *RentHandler) GetRentsByLandlord(ctx context.Context, req *pb.GetRentsByLandlordRequest) (*pb.GetRentsByLandlordResponse, error) {
	rents, err := h.rentUseCase.GetRentsByLandlord(req.LandlordId)
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

	return &pb.GetRentsByLandlordResponse{Rents: pbRents}, nil
}
