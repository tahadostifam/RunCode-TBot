package code_runner_bot

import (
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func HandleStartCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	message := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi Bro!")

	bot.Send(message)
}

func HandleHelpCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI, help_text string) {
	SendTextMessage(update, bot, help_text)
}

func HandleRunCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	msg := update.Message.Text
	msgWordsArr := strings.Fields(msg)

	lang := msgWordsArr[1]
	code_string := msg[7:]

	if !IsValidLanguage(lang) {
		SendTextMessage(update, bot, "درحال حاضر از زبان وارد شده پشتیبانی نمیکنیم!")
		return
	}

	SendTextMessage(update, bot, fmt.Sprintf("lang : %s\ncode : %s\n", lang, strings.TrimSpace(code_string)))
}

func IsValidLanguage(name string) bool {
	for _, l := range []string{"js", "py", "rb"} {
		if l == name {
			return true
		}
	}
	return false
}

func RemoveAllSpacesOfText(text string) string {
	fn_str := ""
	for i := 0; i < len(text); i++ {
		s := text[i]
		if len(strings.TrimSpace(string(s))) > 0 {
			fn_str += string(s)
		}
	}
	return fn_str
}
