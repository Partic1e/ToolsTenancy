package grpc

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "rentservice/api/rent"
	"rentservice/internal/core/usecase"
)

type RentHandler struct {
	pb.UnimplementedRentServiceServer
	getRentsByLandlordUseCase *usecase.GetRentsByLandlordUseCase
	getRentsByRenterUseCase   *usecase.GetRentsByRenterUseCase
	getRentedDatesUseCase     *usecase.GetRentedDatesUseCase
}

func NewRentHandler(
	getRentsByLandlordUseCase *usecase.GetRentsByLandlordUseCase,
	getRentsByRenterUseCase *usecase.GetRentsByRenterUseCase,
	getRentedDatesUseCase *usecase.GetRentedDatesUseCase,
) *RentHandler {
	return &RentHandler{
		getRentsByLandlordUseCase: getRentsByLandlordUseCase,
		getRentsByRenterUseCase:   getRentsByRenterUseCase,
		getRentedDatesUseCase:     getRentedDatesUseCase,
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
			HeldId:     rent.HeldID,
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
			HeldId:     rent.HeldID,
		})
	}

	return &pb.GetResponse{Rents: rentList}, nil
}

func (h *RentHandler) GetRentedDates(ctx context.Context, req *pb.GetRentedDatesRequest) (*pb.GetRentedDatesResponse, error) {
	dates, err := h.getRentedDatesUseCase.GetRentedDates(req.AdId)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении арендованных дат: %v", err)
	}

	var pbDates []*timestamppb.Timestamp
	for _, d := range dates {
		pbDates = append(pbDates, timestamppb.New(d))
	}

	return &pb.GetRentedDatesResponse{RentedDates: pbDates}, nil
}
