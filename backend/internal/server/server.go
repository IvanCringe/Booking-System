package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IvanCringe/Booking-System/backend/internal/config"
)

//* Ответ эндпоинта /health: статус работы сервера
type HealthResponse struct {
	Status string `json:"status"`
}

//* Run регистрирует маршруты и запускает HTTP-сервер
func Run(cfg config.Config) error {

	//* Регистрация маршрутов
	http.HandleFunc("/health", healthHandler)

	//* Корневой маршрут: проверка, что сервер отвечает
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "booking API is running!")
	})

	//* Сообщение в лог о старте сервера
	log.Println("Server started...")
	addr := ":" + cfg.HTTPPort

	//* Адрес сервера из конфигурации
	return http.ListenAndServe(addr, nil)
}

//* healthHandler отвечает JSON-статусом на GET /health
func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := HealthResponse{Status: "ok"}
	json.NewEncoder(w).Encode(s)
}
