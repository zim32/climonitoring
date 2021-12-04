package main

import (
	"bufio"
	"flag"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/zim32/climonitoring/utils"
)

type CliOptions struct {
	ChatID int64
	Token  string
}

func main() {
	options := parseOptions()
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		bot, err := tgbotapi.NewBotAPI(options.Token)
		utils.CatchError(err)

		message := utils.UnMarshalMessage(text)
		telegramMessage := tgbotapi.NewMessage(options.ChatID, "Message: "+message.Message)

		_, err = bot.Send(telegramMessage)
		utils.CatchError(err)

		_, err = os.Stdout.WriteString(text)
		utils.CatchError(err)
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr := utils.GetConfigInt("cid", -1, "Chat ID", "telegram")
	strPtr := utils.GetConfigString("token", "", "Bot token", "telegram")

	flag.Parse()

	options.ChatID = int64(*intPtr)
	options.Token = *strPtr

	if options.ChatID == -1 {
		panic("Chat ID not specified")
	}

	utils.ValidateEmptyString(options.Token, "Api token required")

	return options
}
