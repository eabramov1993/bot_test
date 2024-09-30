package telegram

import (
	"log"
	"os"
	"test_bot/internal/handlers"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var botToken = os.Getenv("TELEGRAM_TOKEN")

// StartBot запускает бота
func StartBot() error {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	// bot.Debug = true // отладчик!!

	log.Printf("Connect %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		// Обработка текстовых сообщений
		if update.Message != nil {
			if update.Message.IsCommand() {
				handlers.HandleCommand(bot, update)
			} else {
				handlers.HandleButton(bot, update)
			}
		}

		// Обработка колбэков
		if update.CallbackQuery != nil {
			handlers.HandleCallback(bot, update)
		}
	}

	return nil
}
