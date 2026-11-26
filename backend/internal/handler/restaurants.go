package handler

import (
	"net/http"
)

//* ListRestaurants возвращает список всех ресторанов (заготовка)
func ListRestaurants(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* GetRestaurant возвращает один ресторан по id (заготовка)
func GetRestaurant(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* CreateRestaurant создает новый ресторан, доступно владельцу (заготовка)
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* UpdateRestaurant изменяет данные ресторана, доступно владельцу (заготовка)
func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}

//* DeleteRestaurant удаляет ресторан, доступно владельцу (заготовка)
func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	writeError(w, http.StatusNotImplemented, "not_implemented", "не реализовано")
}
