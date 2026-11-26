package handler

import (
	"net/http"
)

//* ListBookings возвращает брони текущего пользователя (заготовка)
func ListBookings(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* CreateBooking создает бронь столика (заготовка)
func CreateBooking(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* DeleteBooking отменяет бронь по id (заготовка)
func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* ListRestaurantBookings возвращает брони ресторана, доступно владельцу (заготовка)
func ListRestaurantBookings(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}
