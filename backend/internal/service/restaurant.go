package service

import (
	"context"
	"errors"

	"github.com/IvanCringe/Booking-System/backend/internal/domain"
)

//* RestaurantRepository описывает, что service нужно от хранилища ресторанов (реализует repository)
type RestaurantRepository interface {
	//* GetByID возвращает ресторан по id, при отсутствии domain.ErrNotFound
	GetByID(ctx context.Context, id int64) (*domain.Restaurant, error)
	//* Create сохраняет новый ресторан и возвращает его id
	Create(ctx context.Context, restaurant domain.Restaurant) (int64, error)
	//* List возвращает все рестораны
	List(ctx context.Context) ([]*domain.Restaurant, error)
	//* Update меняет данные ресторана и возвращает обновленный, при отсутствии domain.ErrNotFound
	Update(ctx context.Context, restaurant domain.Restaurant) (*domain.Restaurant, error)
	//* Delete удаляет ресторан по id, при отсутствии domain.ErrNotFound
	Delete(ctx context.Context, id int64) error
}

//* RestaurantService содержит бизнес-логику ресторанов и работает с данными через repository
type RestaurantService struct {
	repo RestaurantRepository
}

//* NewRestaurantService создает сервис ресторанов с переданным repository
func NewRestaurantService(repo RestaurantRepository) *RestaurantService {
	return &RestaurantService{repo: repo}
}

//* ErrInvalidID возвращается, когда id ресторана некорректный
var ErrInvalidID = errors.New("некорректный id")

//* GetByID возвращает ресторан по id, при некорректном id ErrInvalidID
func (s *RestaurantService) GetByID(ctx context.Context, id int64) (*domain.Restaurant, error) {
	if id <= 0 {
		return nil, ErrInvalidID
	}
	return s.repo.GetByID(ctx, id)
}

//* List возвращает все рестораны
func (s *RestaurantService) List(ctx context.Context) ([]*domain.Restaurant, error) {
	return s.repo.List(ctx)
}

//* Delete удаляет ресторан по id, при некорректном id ErrInvalidID
func (s *RestaurantService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidID
	}
	return s.repo.Delete(ctx, id)
}
