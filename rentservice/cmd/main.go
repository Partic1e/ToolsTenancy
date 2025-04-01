package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "rentservice/api/rent"
	"rentservice/config"
	"rentservice/internal/core/usecase"
	grpcd "rentservice/internal/delivery/grpc"
	"rentservice/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[RentService] ❌  - Ошибка загрузки конфига: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+cfg.Server.ServerPort)
	if err != nil {
		log.Fatalf("[RentService][grpc] ❌  - Ошибка при запуске сервера: %v", err)
	}

	db, err := repository.NewDB(&cfg.Database)
	if err != nil {
		log.Fatalf("[RentService][Postgres] ❌  - Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	server := grpc.NewServer()

	rentRepo := repository.NewRentRepository(db)
	getRentsByLandlordUseCase := usecase.NewGetRentsByLandlordUseCase(rentRepo)
	getRentsByRenterUseCase := usecase.NewGetRentsByRenterUseCase(rentRepo)

	rentHandler := grpcd.NewRentHandler(getRentsByLandlordUseCase, getRentsByRenterUseCase)

	pb.RegisterRentServiceServer(server, rentHandler)

	reflection.Register(server)

	log.Println("[RentService][grpc] ✅  - Запущен на :" + cfg.Server.ServerPort)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("[RentService][grpc] ❌  - Ошибка сервера: %v", err)
	}
}
