package service

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/IvanCringe/Booking-System/backend/internal/domain"
)

//* ErrInvalidID возвращается, когда id ресторана некорректный
var ErrInvalidID = errors.New("некорректный id")

//* ErrInvalidRestaurant возвращается, когда данные ресторана не прошли проверку
var ErrInvalidRestaurant = errors.New("некорректные данные ресторана")

//* maxNameLen, maxAddressLen – лимиты длины полей, как в миграции (в символах)
const maxNameLen, maxAddressLen = 150, 255

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

//* Create проверяет данные ресторана и сохраняет его, при некорректных данных ErrInvalidRestaurant
func (s *RestaurantService) Create(ctx context.Context, restaurant domain.Restaurant) (int64, error) {
	if err := validateRestaurant(restaurant); err != nil {
		return 0, err
	}
	return s.repo.Create(ctx, restaurant)
}

//* Delete удаляет ресторан по id, при некорректном id ErrInvalidID
func (s *RestaurantService) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return ErrInvalidID
	}
	return s.repo.Delete(ctx, id)
}

//* isValidPhone проверяет российский номер: 8XXXXXXXXXX или +7XXXXXXXXXX
func isValidPhone(phone string) bool {
	var r string
	switch {
	case strings.HasPrefix(phone, "+7"):
		r = phone[2:]
	case strings.HasPrefix(phone, "8"):
		r = phone[1:]
	default:
		return false
	}
	if len(r) != 10 {
		return false
	}
	_, err := strconv.ParseUint(r, 10, 64)
	if err != nil {
		return false
	}
	return true
}

//* validateRestaurant проверяет данные ресторана, при ошибке возвращает ErrInvalidRestaurant с пояснением
func validateRestaurant(r domain.Restaurant) error {
	if strings.TrimSpace(r.Name) == "" {
		return fmt.Errorf("%w: имя пустое", ErrInvalidRestaurant)
	}
	if utf8.RuneCountInString(r.Name) > maxNameLen {
		return fmt.Errorf("%w: имя длиннее %d символов", ErrInvalidRestaurant, maxNameLen)
	}
	if strings.TrimSpace(r.Address) == "" {
		return fmt.Errorf("%w: адрес пустой", ErrInvalidRestaurant)
	}
	if utf8.RuneCountInString(r.Address) > maxAddressLen {
		return fmt.Errorf("%w: адрес длиннее %d символов", ErrInvalidRestaurant, maxAddressLen)
	}
	if !isValidPhone(r.Phone) {
		return fmt.Errorf("%w: некорректный формат телефона", ErrInvalidRestaurant)
	}
	if r.OwnerId <= 0 {
		return fmt.Errorf("%w: некорректный владелец", ErrInvalidRestaurant)
	}
	return nil
}
