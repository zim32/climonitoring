package main

import (
	"climonitoring/utils"
	"bufio"
	"encoding/json"
	"flag"
	"log"
	"os"
	"strings"
	"time"
)

type CliOptions struct {
	Message  string
	Severity string
	HostName string
}


func main() {
	reader  := bufio.NewReader(os.Stdin)
	options := parseOptions()

	for {
		text, err := utils.GetNewLine(reader)

		var messageText string

		if len(options.Message) > 0 {
			messageText = options.Message
		} else {
			messageText = strings.Trim(text, utils.EOT_S)
		}

		messageText = strings.Replace(messageText, "{stdin}", strings.Trim(text, utils.EOT_S), -1)

		msg := utils.NewMessage()
		msg.Severity = options.Severity
		msg.Message  = messageText
		msg.Created  = time.Now()

		if len(options.HostName) > 0 {
			msg.HostName = options.HostName
		}

		b, err := json.Marshal(msg)

		_, err = os.Stdout.WriteString(string(b) + utils.EOT_S)

		if err != nil {
			log.Print(err)
			os.Exit(0)
		}
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	strPtrM := utils.GetConfigString("m", "", "Message", "message")
	strPtrS := utils.GetConfigString("s", "info", "Severity", "message")
	strPtrH := utils.GetConfigString("h", "", "Host name", "message")

	flag.Parse()

	options.Message  = *strPtrM
	options.Severity = *strPtrS
	options.HostName = *strPtrH

	return options
}