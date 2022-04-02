package code_runner_bot

import (
	"fmt"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var help_text string

func RunBot(bot *tgbotapi.BotAPI) {
	FetchAndSetHelpCommandText()
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

		// Definfing Commands
		var recvd_msg string = strings.TrimSpace(update.Message.Text)
		is_bad_command := true
		switch recvd_msg {
		case "/start":
			is_bad_command = false
			HandleStartCommand(&update, bot)
		case "/help":
			is_bad_command = false
			HandleHelpCommand(&update, bot, help_text)
		}
		if recvd_msg == "/run" {
			is_bad_command = false

			var run_help string
			run_help += "برای اجرای کد تان این کامند را اجرا کنید :\n"
			run_help += "/run <js,py,rb> Your-Code"
			run_help += "\n\n"
			run_help += "توجه کنید دسترسی به اینترنت مجاز نمی باشد و ران کردن هرگونه برنامه سنگین جهت *آیش سرور ما با بن شدن شما طرف خواهد شد پس سعی کنید از while و یا for بی نهایت استفاده نکنید."

			SendTextMessage(&update, bot, run_help)
		} else if strings.Contains(recvd_msg, "/run") {
			is_bad_command = false

			HandleRunCommand(&update, bot)
		}
		if is_bad_command {
			HandleUndefinedCommand(&update, bot)
		}
		// Definfing Commands End

		// FIXME - just devel
		fmt.Printf("Message ID: %d\n", update.Message.Chat.ID)
		fmt.Printf("Message TEXT: %s\n", update.Message.Text)
		fmt.Println("--------------------")
	}
}

func FetchAndSetHelpCommandText() {
	cwd, cwd_err := os.Getwd()
	if cwd_err != nil {
		fmt.Println("An error occurred on getting cwd")
		os.Exit(1)
	} else {
		file, file_err := os.ReadFile(cwd + "/help_command.help.txt")
		if file_err != nil {
			fmt.Printf("An error occurred on reading help_command.txt file from -> %s\n", cwd)
			os.Exit(1)
		}
		help_text = string(file)
	}
}

func HandleUndefinedCommand(update *tgbotapi.Update, bot *tgbotapi.BotAPI) {
	SendTextMessage(update, bot, "کامند وارد شده صحیح نمی باشد خواهشا گمشید خودتون رو با /help از جهالت در بیاورید! تامام.")
}

func SendTextMessage(update *tgbotapi.Update, bot *tgbotapi.BotAPI, message string) {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)

	bot.Send(msg)
}
