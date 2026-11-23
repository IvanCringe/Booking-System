package handler

import (
	"net/http"
)

//* AuthRegistration регистрирует нового пользователя (заготовка)
func AuthRegistration(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* AuthLogin выполняет вход и выдает токен (заготовка)
func AuthLogin(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}
