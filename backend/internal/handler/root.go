package handler

import (
	"fmt"
	"net/http"
)

//* RootHandler отвечает текстом, что сервер работает
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "booking API is running!")
}
