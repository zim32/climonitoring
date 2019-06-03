package main

import (
	"bufio"
	"climonitoring/utils"
	"flag"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

type CliOptions struct {
	ChatID int64
	Token  string
}


func main() {
	options := parseOptions()
	reader  := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		bot, err := tgbotapi.NewBotAPI(options.Token)
		utils.CatchError(err)

		message         := utils.UnMarshalMessage(text)
		telegramMessage := tgbotapi.NewMessage(options.ChatID, "Message: " + message.Message)

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
	options.Token  = *strPtr

	if options.ChatID == -1 {
		panic("Chat ID not specified")
	}

	if len(options.Token) == 0 {
		panic("Token not specified")
	}

	return options
}