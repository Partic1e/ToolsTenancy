package main

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"
	"mainservice/internal/delivery/httpengine"
	"mainservice/internal/delivery/httpengine/handler"
	"mainservice/internal/lib/grpcclient"
)

type Config struct {
	GRPC grpcclient.GrpcConfig `mapstructure:"GRPC"`
	HTTP httpengine.HttpConfig `mapstructure:"HTTP"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("error reading cfg: %v", err)
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("error unmarshal cfg: %v", err)
		return nil, err
	}

	return &cfg, nil
}

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	userClient := grpcclient.NewGrpcClient(cfg.GRPC.UserService.Host, cfg.GRPC.UserService.Port)
	paymentClient := grpcclient.NewGrpcClient(cfg.GRPC.PaymentService.Host, cfg.GRPC.PaymentService.Port)
	adClient := grpcclient.NewGrpcClient(cfg.GRPC.AdService.Host, cfg.GRPC.AdService.Port)
	rentClient := grpcclient.NewGrpcClient(cfg.GRPC.RentService.Host, cfg.GRPC.RentService.Port)

	h := handler.NewHandler(userClient, paymentClient, adClient, rentClient)
	r := httpengine.InitRouter(context.Background(), h)

	addr := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	log.Printf("start server on %s", addr)
	log.Fatal(r.Run(addr))
}
