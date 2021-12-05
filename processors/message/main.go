package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/zim32/climonitoring/utils"
)

type CliOptions struct {
	Message         string
	Severity        string
	HostName        string
	MessagesPerHour int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	options := parseOptions()
	hourCounter := time.Now()
	messagesCounter := int64(0)
	limitReached := false

	for {
		text, err := utils.GetNewLine(reader)

		// reset messages counter every hour
		if time.Now().Sub(hourCounter).Hours() > 1 {
			messagesCounter = 0
			hourCounter = time.Now()
			limitReached = false
		}

		if limitReached {
			continue
		}

		if messagesCounter > int64(options.MessagesPerHour) {
			limitReached = true
			text = fmt.Sprintf("More than %d messages per hour. Skipping new messages...", options.MessagesPerHour) + utils.EOT_S
		}

		var messageText string

		if len(options.Message) > 0 {
			messageText = options.Message
		} else {
			messageText = strings.Trim(text, utils.EOT_S)
		}

		messageText = strings.Replace(messageText, "{stdin}", strings.Trim(text, utils.EOT_S), -1)

		if len(messageText) == 0 {
			continue
		}

		msg := utils.NewMessage()
		msg.Severity = options.Severity
		msg.Message = messageText
		msg.Created = time.Now()

		if len(options.HostName) > 0 {
			msg.HostName = options.HostName
		}

		b, err := json.Marshal(msg)

		_, err = os.Stdout.WriteString(string(b) + utils.EOT_S)
		utils.CatchError(err)

		messagesCounter++
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	strPtrM := utils.GetConfigString("m", "", "Message", "message")
	strPtrS := utils.GetConfigString("s", "info", "Severity", "message")
	strPtrH := utils.GetConfigString("h", "", "Host name", "message")
	intPtr1 := utils.GetConfigInt("l", 60, "Max messages per hour", "message")

	flag.Parse()

	options.Message = *strPtrM
	options.Severity = *strPtrS
	options.HostName = *strPtrH
	options.MessagesPerHour = *intPtr1

	return options
}
