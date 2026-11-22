package server

import (
	"net/http"

	"github.com/IvanCringe/Booking-System/backend/internal/handler"
)

func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", handler.HealthHandler)
	mux.HandleFunc("/{$}", handler.RootHandler)
	return mux
}
