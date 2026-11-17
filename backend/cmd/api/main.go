package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/IvanCringe/Booking-System/backend/internal/config"
	"github.com/IvanCringe/Booking-System/backend/internal/database"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func main() {
	//* Создаем конфиг
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	//* Создаем пул
	pool, err := database.NewPostgresPool(cfg)
	if err != nil {
		log.Fatalf("failed to open pool: %v", err)
	}
	defer pool.Close()

	http.HandleFunc("/health", healthHandler)
	//* Запускаем HTTP-сервер и обрабатываем ошибки.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "booking API is running!")
	})

	log.Println("Server started on :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := HealthResponse{Status: "ok"}
	json.NewEncoder(w).Encode(s)
}
