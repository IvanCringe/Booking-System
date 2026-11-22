package handler

import (
	"encoding/json"
	"net/http"
)

//* Ответ эндпоинта /health: статус работы сервера
type HealthResponse struct {
	Status string `json:"status"`
}

//* HealthHandler отвечает JSON-статусом на GET /health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	//* Сообщаем клиенту, что в ответе JSON
	w.Header().Set("Content-Type", "application/json")
	//* Формируем и отправляем ответ
	s := HealthResponse{Status: "ok"}
	json.NewEncoder(w).Encode(s)
}
