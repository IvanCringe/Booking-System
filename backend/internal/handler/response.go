package handler

import (
	"encoding/json"
	"net/http"
)

//* ErrorResponse описывает ответ с ошибкой: текстовый код для программы и сообщение для человека
type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

//* writeJSON записывает ответ в формате JSON с указанным HTTP-статусом
func writeJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	e := ErrorResponse{Code: code, Message: message}
	//* writeError отправляет клиенту ошибку в едином формате
	writeJSON(w, status, e)
}
