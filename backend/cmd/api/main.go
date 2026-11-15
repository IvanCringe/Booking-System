package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/IvanCringe/Booking-System/backend/internal/config"
    "github.com/IvanCringe/Booking-System/backend/internal/database"
)

func main() {
    //* Создаем конфиг
    cfg, err := config.LoadConfig()
    if err != nil {
        log.Fatalf("failed to load config: %v", err)
    }
    
    //* Создаем пул 
    pool, err := database.NewPostgresPool(cfg)
    if err != nil {
        log.Fatalf("failed to open pool: %v", err)
    }
    defer pool.Close()

    //* Запускаем HTTP-сервер и обрабатываем ошибки.
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintln(w, "booking API is running!")
    })

    log.Println("Server started on :8080")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}