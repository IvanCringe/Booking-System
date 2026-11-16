package database

import (
    "context"
    "fmt"
    "time"

    "github.com/jackc/pgx/v5/pgxpool"
    "github.com/IvanCringe/Booking-System/backend/internal/config"
)

func NewPostgresPool(cfg config.Config) (*pgxpool.Pool, error) {
    //* Создаём контекст с ограничением времени в 5 секунд. 
    ctx, cancel := context.WithTimeout(
        context.Background(),
        5*time.Second,
    )
    
    //? Завершает контекст 
    defer cancel()

    //* Собираем строку подключения (данные беруться из config/config.go)
    connString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword,cfg.DBName)

    //* Создаём объект пула соединений
    pool, err := pgxpool.New(ctx, connString)
    if err != nil {
        return nil, fmt.Errorf("create postgres pool: %w", err)
    }

    //* Проверка подключения к бд
    if err := pool.Ping(ctx); err != nil {
        pool.Close()
        return nil, fmt.Errorf("ping postgres: %w", err)
    }
    return pool, nil
}