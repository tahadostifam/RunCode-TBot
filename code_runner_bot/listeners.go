package code_runner_bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStartCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi Bro!")

	bot.Send(message)
}
