package code_runner_bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func RunBot(bot *tgbotapi.BotAPI) {
	HandleUpdates(bot)
}

func HandleUpdates(bot *tgbotapi.BotAPI) {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := bot.GetUpdatesChan(updateConfig)
	for update := range updates {
		if update.Message == nil {
			continue
		}

		switch strings.TrimSpace(update.Message.Text) {
		case "/start":
			HandleStartCommand(&update, bot)
		default:
			HandleUndefinedCommand(&update, bot)
		}

		// FIXME - just devel
		fmt.Printf("Message ID: %d\n", update.Message.Chat.ID)
		fmt.Printf("Message TEXT: %s\n", update.Message.Text)
		fmt.Println("--------------------")
	}
}

func HandleUndefinedCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	SendTextMessage(update, bot, "Command Not Found! Please use /help to get out of ignorance :]")
}

func SendTextMessage(update *tgbotapi.Update, bot *tgbotapi.BotAPI, message string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	bot.Send(msg)
}
