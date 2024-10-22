package handlers

import (
	"context"
	"fmt"
	"log"
	"strings"
	"test_bot/internal/scraper"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const red = "\033[31m"
const reset = "\033[0m"

var allowedUsers = map[int64]bool{
	1924014411: true,
	816116066:  true,
}

// var allowedCommands = map[string]bool{
// 	"/start": true,
// 	"/help":  true,
// }

// HandleCommand обрабатывает команды
func HandleCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	// Проверяем, есть ли пользователь в списке разрешённых
	if !allowedUsers[chatID] {
		log.Printf("%sНовый пользователь: %d запрашивает доступ!%s", red, chatID, reset)

		msg := tgbotapi.NewMessage(chatID, "У вас нет доступа к этому боту. Запросите доступ у администратора.")
		bot.Send(msg)
		return
	}

	switch update.Message.Command() {
	case "start":
		keyboard := tgbotapi.NewReplyKeyboard(
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Получить данные BCC"),
				tgbotapi.NewKeyboardButton("Получить данные Kaspi"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Получить данные Халык"),
				tgbotapi.NewKeyboardButton("Получить данные РБК"),
			),
			tgbotapi.NewKeyboardButtonRow(
				tgbotapi.NewKeyboardButton("Получить данные Jusan"),
				tgbotapi.NewKeyboardButton("Получить данные Bereke"),
			),
		)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
		msg.ReplyMarkup = keyboard
		bot.Send(msg)

	default:
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Я не знаю такую команду.")
		bot.Send(msg)
	}
}

// HandleButton обрабатывает нажатия кнопок
func HandleButton(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	source := ""
	switch update.Message.Text {
	case "Получить данные BCC":
		source = "bcc"
	case "Получить данные Халык":
		source = "halyk"
	case "Получить данные Kaspi":
		source = "kaspi"
	case "Получить данные Jusan":
		source = "jusan"
	case "Получить данные Bereke":
		source = "bereke"
	case "Получить данные РБК":
		source = "rbk"
	default:
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	results, err := scraper.ScrapeData(ctx, source)
	if err != nil {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Ошибка при парсинге данных.")
		bot.Send(msg)
		return
	}

	// Отправляем каждую карточку в отдельном сообщении
	for _, result := range results {
		response := fmt.Sprintf("Карта: %s\nСумма: %s", result.CardNumber, result.Amount)

		// Создаем кнопку с передачей номера карты и суммы в CallbackData
		button := tgbotapi.NewInlineKeyboardButtonData("Оплатить", fmt.Sprintf("pay_%s_%s", result.CardNumber, result.Amount))
		keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyMarkup = &keyboard

		bot.Send(msg)
	}
}

// Обработчик колбэков
func HandleCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	if update.CallbackQuery != nil {
		// Проверяем, начинается ли CallbackData с "pay_"
		if strings.HasPrefix(update.CallbackQuery.Data, "pay_") {
			// Извлекаем данные из CallbackData
			dataParts := strings.Split(strings.TrimPrefix(update.CallbackQuery.Data, "pay_"), "_")
			if len(dataParts) != 2 {
				log.Println("Неверный формат данных в CallbackData")
				return
			}
			cardNumber := dataParts[0]
			amount := dataParts[1]

			// Формируем сообщение об успешной оплате
			response := fmt.Sprintf("✅Карта: %s\n✅Сумма: %s", cardNumber, amount)

			// Создаем новую инлайн-кнопку с текстом "Оплачено"
			button := tgbotapi.NewInlineKeyboardButtonData("Оплачено ✅", "payment_done")
			keyboard := tgbotapi.NewInlineKeyboardMarkup(tgbotapi.NewInlineKeyboardRow(button))

			// Редактируем текущее сообщение, чтобы изменить текст и клавиатуру
			editMsg := tgbotapi.NewEditMessageText(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Message.MessageID, response)
			editMsg.ReplyMarkup = &keyboard // Передаем указатель на клавиатуру

			// Отправляем редактированное сообщение
			if _, err := bot.Send(editMsg); err != nil {
				log.Printf("Error sending edited message: %v", err)
			}

			// // Отправляем ответ на колбэк с сообщением
			// msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Оплата успешна!")
			// if _, err := bot.Send(msg); err != nil {
			// 	log.Printf("Error sending message after callback: %v", err)
			// }
		}
	}
}
