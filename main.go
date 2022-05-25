package main

import (
	"CODE-Runner/code_runner_bot"
	"CODE-Runner/configs"
	"fmt"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var token string

func main() {

	configs.ReadAndSetConfigs()
	code_runner_bot.SetSSHClient()

	token = configs.AllConfigs.BOT_TOKEN

	botClient, botErr := tgbotapi.NewBotAPI(token)
	if botErr != nil {
		fmt.Printf("Error: %v\n", botErr)
		os.Exit(1)
	}

	fmt.Println("Welcome to CODE-Runner Bot")

	code_runner_bot.RunBot(botClient)

}
