package config

import (
	"fmt"
	"os"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	HTTPPort   string
}

func LoadConfig() (Config, error) {
	//* Запись значений из окружения
	cfg := Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		HTTPPort:   os.Getenv("HTTP_PORT"),
	}

	//* Проверка наличия данных окружения
	if cfg.DBHost == "" {
		return Config{}, fmt.Errorf("DB_HOST is required")
	}
	if cfg.DBPort == "" {
		return Config{}, fmt.Errorf("DB_PORT is required")
	}
	if cfg.DBUser == "" {
		return Config{}, fmt.Errorf("DB_USER is required")
	}
	if cfg.DBPassword == "" {
		return Config{}, fmt.Errorf("DB_PASSWORD is required")
	}
	if cfg.DBName == "" {
		return Config{}, fmt.Errorf("DB_NAME is required")
	}
	if cfg.HTTPPort == "" {
		cfg.HTTPPort = "8080"
	}
	return cfg, nil
}
