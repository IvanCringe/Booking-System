package server

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IvanCringe/Booking-System/backend/internal/config"
)

//* Run запускает HTTP-сервер и останавливает его аккуратно по сигналу
func Run(cfg config.Config) error {

	//* Сообщение в лог о старте сервера
	log.Println("Server started...")

	//* Адрес сервера из конфигурации
	addr := ":" + cfg.HTTPPort

	//* Свой сервер в переменной, чтобы можно было вызвать у него Shutdown
	srv := &http.Server{Addr: addr, Handler: newRouter()}

	//* Сервер работает в отдельной горутине, чтобы основной поток мог ждать сигнал
	go func() {
		err := srv.ListenAndServe()
		//* ErrServerClosed приходит при штатной остановке, это не ошибка
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Printf("ошибка сервера %v", err)
		}
	}()

	//* Контекст, который отменяется при Ctrl+C или SIGTERM
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	//* Ждем сигнал остановки
	<-ctx.Done()

	log.Printf("останавливаю сервер...")

	//* Новый контекст: ctx уже отменен, а Shutdown нужно время, чтобы дождаться текущих запросов
	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	//* Аккуратная остановка: новые запросы не принимаем, текущие доделываем
	err := srv.Shutdown(shutdownCtx)
	if err != nil {
		log.Printf("ошибка при остановке сервера: %v", err)
		return err
	}

	log.Printf("сервер остановлен")
	return nil
}
