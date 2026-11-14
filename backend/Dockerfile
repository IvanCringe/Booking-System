# ==========================================
# STAGE 1: BUILD
# ==========================================
# Используем легковесный alpine образ Go для сборки
FROM golang:1.26-alpine AS builder
WORKDIR /app

# Устанавливаем git и сертификаты (часто нужны для загрузки приватных Go-модулей)
RUN apk add --no-cache git ca-certificates

# Копируем файлы зависимостей
COPY go.mod go.sum ./

# Загружаем зависимости
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем приложение с оптимизацией размера бинарника (убираем отладочную информацию)
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
    -ldflags="-s -w" \
    -o /app/bin/api ./cmd/api


# ==========================================
# STAGE 2: RUN
# ==========================================
FROM alpine:3.20

WORKDIR /app

# Добавляем часовые пояса и сертификаты для работы с HTTPS/внешними API
RUN apk add --no-cache tzdata ca-certificates

# Копируем только готовый бинарник
COPY --from=builder /app/bin/api ./api

# Создаем безопасного пользователя (не запускаем процесс от root)
RUN adduser -D -u 10001 appuser
USER appuser

EXPOSE 8080

CMD ["./api"]
