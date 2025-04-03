package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	Server   Server         `mapstructure:"grpc"`
	Database DatabaseConfig `yaml:"database"`
}

type Server struct {
	ServerPort string `mapstructure:"server_port"`
}

type DatabaseConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Ошибка чтения конфигурационного файла: %v", err)
		return nil, err
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Ошибка разбора конфига: %v", err)
		return nil, err
	}

	return &config, nil
}
