package server

import (
	"net/http"

	"github.com/IvanCringe/Booking-System/backend/internal/handler"
)

//* newRouter создает роутер и регистрирует все маршруты API
func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	//* Проверка состояния сервера
	mux.HandleFunc("/health", handler.HealthHandler)
	//* Корневой маршрут, только точное совпадение пути /
	mux.HandleFunc("/{$}", handler.RootHandler)
	return mux
}
