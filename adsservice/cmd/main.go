package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "adsservice/api/ad"
	"adsservice/config"
	"adsservice/internal/core/usecase"
	grpcd "adsservice/internal/delivery/grpc"
	"adsservice/internal/repository"
	"google.golang.org/grpc"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[AdsService] ❌  - Ошибка загрузки конфига: %v", err)
	}

	listener, err := net.Listen("tcp", ":"+cfg.Server.ServerPort)
	if err != nil {
		log.Fatalf("[AdsService][grpc] ❌  - Ошибка при запуске сервера: %v", err)
	}

	db, err := repository.NewDB(&cfg.Database)
	if err != nil {
		log.Fatalf("[AdsService][Postgres] ❌  - Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	adRepo := repository.NewAdRepository(db)
	adUseCase := usecase.NewAdUseCase(adRepo)
	adHandler := grpcd.NewAdHandler(adUseCase)

	server := grpc.NewServer()
	pb.RegisterAdServiceServer(server, adHandler)

	reflection.Register(server)

	log.Println("[AdsService][grpc] ✅  - Запущен на :" + cfg.Server.ServerPort)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("[AdsService][grpc] ❌  - Ошибка сервера: %v", err)
	}
}
