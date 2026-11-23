package handler

import (
	"net/http"
)

//* GetCurrentUser возвращает данные текущего пользователя (заготовка)
func GetCurrentUser(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* GetUser возвращает данные пользователя по id (заготовка)
func GetUser(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}
