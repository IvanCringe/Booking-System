package domain

import (
	"time"
)

// 3. В структуре у каждого поля указываем Имя и Тип
type RestaurantTable struct {
	ID               int64
	RestaurantID     int64
	Name             string 
	SeatsCount       int64 
	IsActive         bool 
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
