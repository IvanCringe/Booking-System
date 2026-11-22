package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HealthResponse struct {
	Status string `json:"status"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := HealthResponse{Status: "ok"}
	json.NewEncoder(w).Encode(s)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "booking API is running!")
}
