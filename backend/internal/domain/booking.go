package domain

import (
	"time"
)

// 1. Сначала объявляем сам тип Role на базе строки
type Status string

// 2. Затем создаём константы этого типа
const (
	StatusPending    Status = "PENDING"
	StatusConfirmed  Status = "CONFIRMED"
	StatusCancelled  Status = "CANCELLED"
	StatusCompleted  Status = "COMPLETED"
)

// 3. В структуре у каждого поля указываем Имя и Тип
type Booking struct {
	ID             int64
	UserID         int64
	TableID        int64
	GuestCount     int64
	StartTime      time.Time
	EndTime        time.Time
	Status         Status
	Comment        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

