package main

import (
	"log"
	"test_bot/internal/telegram"
)

func main() {
	err := telegram.StartBot()
	if err != nil {
		log.Fatalf("Ошибка запуска бота: %v", err)
	}
}
