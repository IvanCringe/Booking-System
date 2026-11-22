package server

import (
	"log"
	"net/http"

	"github.com/IvanCringe/Booking-System/backend/internal/config"
)

//* Run регистрирует маршруты и запускает HTTP-сервер
func Run(cfg config.Config) error {

	//* Сообщение в лог о старте сервера
	log.Println("Server started...")

	//* Адрес сервера из конфигурации
	addr := ":" + cfg.HTTPPort

	return http.ListenAndServe(addr, newRouter())
}
