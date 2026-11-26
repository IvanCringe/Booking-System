package domain

import (
	"time"
)

// 1. Сначала объявляем сам тип Role на базе строки
type Role string

// 2. Затем создаём константы этого типа
const (
	RoleAdmin  Role = "ADMIN"
	RoleOwner  Role = "OWNER"
	RoleClient Role = "CLIENT"
)

// 3. В структуре у каждого поля указываем Имя и Тип
type User struct {
	ID        int64
	Name      string
	Email     string    
	CreatedAt time.Time 
	Password  string    
	Role      Role      
}
