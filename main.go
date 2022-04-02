package main

import (
	"CODE-Runner/code_runner_bot"
	"CODE-Runner/configs"
	"fmt"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token string = strings.TrimSpace(os.Getenv("CODE_RUNNER_BOT_TOKEN"))

func main() {
	if len(token) > 0 {
		configs.ReadAndSetConfigs()
		code_runner_bot.SetSSHClient()

		bot_client, bot_err := tgbotapi.NewBotAPI(token)
		if bot_err != nil {
			fmt.Printf("Error: %v\n", bot_err)
			os.Exit(1)
		}

		fmt.Println("Welcome to CODE-Runner Bot")

		code_runner_bot.RunBot(bot_client)

	} else {
		fmt.Println("Bot token not found. Please set token using :\nexport CODE_RUNNER_BOT_TOKEN=your-token")
	}
}
