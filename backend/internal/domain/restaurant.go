package domain

import (
	"time"
)

// 3. В структуре у каждого поля указываем Имя и Тип
type Restaurant struct {
	ID          int64
	OwnerId     int64
	Name        string 
	Address     string 
	Phone       string 
	CreatedAt   time.Time
	OpenTime    time.Time
	CloseTime   time.Time
}
