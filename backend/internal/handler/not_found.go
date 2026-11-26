package handler

import (
	"net/http"
)

//* NotFoundHandler отвечает ошибкой в едином формате, если путь не найден
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotFound, "not_found", "страница не найдена")
}
