package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	pb "paymentservice/api/payment"
	"paymentservice/internal/data/db"
	"paymentservice/internal/data/repository"
	"paymentservice/internal/domain/usecases"
	"paymentservice/internal/lib/grpcclient"
	"paymentservice/internal/lib/rabbitclient"
	"paymentservice/internal/transport/handlers"
)

type Config struct {
	GRPC     grpcclient.GrpcConfig       `mapstructure:"grpc"`
	Database db.DatabaseConfig           `mapstructure:"database"`
	RabbitMQ rabbitclient.RabbitMQConfig `mapstructure:"rabbitmq"`
}

func LoadConfig(path string) (*Config, error) {
	viper.SetConfigFile(path)
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурации: %v", err)
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Ошибка разбора конфигурации: %v", err)
		return nil, err
	}

	return &cfg, nil
}

func main() {
	cfg, err := LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	database, err := db.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	orm, err := database.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer orm.Close()

	rabbitURL := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/",
		cfg.RabbitMQ.Login,
		cfg.RabbitMQ.Password,
		cfg.RabbitMQ.Host,
		cfg.RabbitMQ.Port,
	)

	rabbit, err := rabbitclient.NewRabbitClient(rabbitURL, cfg.RabbitMQ.Channel)
	if err != nil {
		log.Fatal(err)
	}
	defer rabbit.Close()

	paymentRepository := repository.NewPaymentRepository(orm)
	useCases := usecases.NewPaymentUseCase(*paymentRepository)
	paymentHandler := handler.NewPaymentHandler(*useCases)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer(grpc.Creds(insecure.NewCredentials()))
	pb.RegisterPaymentServiceServer(grpcServer, paymentHandler)

	reflection.Register(grpcServer)

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
