package main

import (
	"log"
	"net/http"
	"os"
	"test_bot/internal/telegram"
)

func main() {
	// Запускаем бота в отдельной горутине
	go func() {
		err := telegram.StartBot()
		if err != nil {
			log.Fatalf("Ошибка запуска бота: %v", err)
		}
	}()

	// Получаем порт из переменной окружения
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable not set")
	}

	// Запускаем HTTP-сервер для Heroku
	log.Printf("Server is running on port %s", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatalf("Ошибка запуска HTTP сервера: %v", err)
	}
}
