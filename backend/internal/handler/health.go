package handler

import (
	"net/http"
)

//* Ответ эндпоинта /health: статус работы сервера
type HealthResponse struct {
	Status string `json:"status"`
}

//* HealthHandler отвечает JSON-статусом на GET /health
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	//* Формируем и отправляем ответ
	s := HealthResponse{Status: "ok"}
	writeJSON(w, http.StatusOK, s)
}
