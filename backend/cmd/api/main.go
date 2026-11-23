package main

import (
	"log"

	"github.com/IvanCringe/Booking-System/backend/internal/config"
	"github.com/IvanCringe/Booking-System/backend/internal/database"
	"github.com/IvanCringe/Booking-System/backend/internal/server"
)

func main() {
	//* Создаем конфиг
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	//* Создаем пул
	pool, err := database.NewPostgresPool(cfg)
	//* Без базы сервер все равно запускается, ошибка пишется в лог
	if err != nil {
		log.Printf("failed to open pool: %v", err)
	}
	//* Закрываем пул при остановке, только если он был создан
	if pool != nil {
		defer pool.Close()
	}

	//* Запускаем HTTP-сервер
	if err := server.Run(cfg); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
