package domain

import (
	"errors"
)

// Заранее заводим понятные ошибки, чтобы не тащить ошибки БД в бизнес-логику
var (
	ErrNotFound     = errors.New("record not found")
	ErrInvalidRole  = errors.New("invalid user role")
	ErrAlreadyExist = errors.New("record already exists")
)