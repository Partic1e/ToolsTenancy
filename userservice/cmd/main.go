package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "userservice/api/user"
	"userservice/config"
	"userservice/internal/core/usecase"
	grpcd "userservice/internal/delivery/grpc"
	"userservice/internal/repository"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("[UserService] ❌  - Ошибка загрузки конфига: %v", err)
	}

	db, err := repository.NewDB(&cfg.Database)
	if err != nil {
		log.Fatalf("[UserService][Postgres] ❌  - Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	deps := usecase.NewDependencies(repository.NewUserRepository(db))
	userHandler := grpcd.NewUserHandler(deps.UserUsecase)

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, userHandler)

	reflection.Register(server)

	listener, err := net.Listen("tcp", ":"+cfg.Server.ServerPort)
	if err != nil {
		log.Fatalf("[UserService][grpc] ❌  - Ошибка запуска сервера: %v", err)
	}

	log.Printf("[UserService][grpc] ✅  - Запущен на :%s\n", cfg.Server.ServerPort)
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Ошибка сервера: %v", err)
	}
}
