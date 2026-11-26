package server

import (
	"net/http"

	"github.com/IvanCringe/Booking-System/backend/internal/handler"
)

// * newRouter создает роутер и регистрирует все маршруты API
func newRouter() *http.ServeMux {
	mux := http.NewServeMux()
	//* Проверка состояния сервера
	mux.HandleFunc("/health", handler.HealthHandler)
	//* Корневой маршрут, только точное совпадение пути /
	mux.HandleFunc("/{$}", handler.RootHandler)
	//* Все остальные пути: ответ 404 в едином формате
	mux.HandleFunc("/", handler.NotFoundHandler)
	//* Рестораны: список, получение, создание, изменение, удаление
	mux.HandleFunc("GET /restaurants", handler.ListRestaurants)
	mux.HandleFunc("GET /restaurants/{id}", handler.GetRestaurant)
	mux.HandleFunc("POST /restaurants", handler.CreateRestaurant)
	mux.HandleFunc("PUT /restaurants/{id}", handler.UpdateRestaurant)
	mux.HandleFunc("DELETE /restaurants/{id}", handler.DeleteRestaurant)
	//* Авторизация: регистрация и вход
	mux.HandleFunc("POST /auth/register", handler.AuthRegistration)
	mux.HandleFunc("POST /auth/login", handler.AuthLogin)
	//* Пользователи: текущий пользователь и по id
	mux.HandleFunc("GET /users/me", handler.GetCurrentUser)
	mux.HandleFunc("GET /users/{id}", handler.GetUser)
	//* Бронирования: свои брони, создание, отмена, брони ресторана
	mux.HandleFunc("GET /bookings", handler.ListBookings)
	mux.HandleFunc("POST /bookings", handler.CreateBooking)
	mux.HandleFunc("DELETE /bookings/{id}", handler.DeleteBooking)
	mux.HandleFunc("GET /restaurants/{id}/bookings", handler.ListRestaurantBookings)

	return mux
}
