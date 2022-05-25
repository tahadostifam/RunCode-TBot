package code_runner_bot

import (
	"fmt"
	"math/rand"
	"os"
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
	codeString := msg[7:]

	if !IsValidLanguage(lang) {
		SendTextMessage(update, bot, "درحال حاضر از زبان وارد شده پشتیبانی نمیکنیم!")
		return
	}

	successExec, codeErr, codeResult := ExecCode(lang, codeString)
	if !successExec {
		SendTextMessage(update, bot, "خطا هنگام اجرای کد :\n"+codeErr)
	} else {
		SendTextMessage(update, bot, fmt.Sprintf("Result: \n%s", codeResult))
	}
}

func ExecCode(lang string, code string) (bool, string, string) {
	randName := rand.Intn(20)
	fileFullPath := fmt.Sprintf("/tmp/%v.code_runner_bot.txt", randName)
	writeFileErr := os.WriteFile(fileFullPath, []byte(code), 0644)
	if writeFileErr != nil {
		os.Remove(fileFullPath)
		return false, "Server Error: Error in writing file on /tmp", ""
	} else {
		var formattedLangCli string
		switch lang {
		case "rb":
			formattedLangCli = fmt.Sprintf("ruby %s", fileFullPath)
		case "py":
			formattedLangCli = fmt.Sprintf("python %s", fileFullPath)
		}

		output, exec_err := sshClient.Cmd(formattedLangCli).Output()

		os.Remove(fileFullPath)

		if exec_err != nil {
			return false, "Error in executing the file: " + exec_err.Error(), ""
		} else {
			return true, "", string(output)
		}
	}
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
