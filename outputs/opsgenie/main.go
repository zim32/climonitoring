package main

import (
	"bufio"
	"bytes"
	"climonitoring/utils"
	"encoding/json"
	"flag"
	"net/http"
	"os"
)

type CliOptions struct {
	ApiToken      string
	ApiEndpoint   string
	ResponderType string
	ResponderId   string
}

type OpsGenieMessage struct {
	Message     string 						`json:"message"`
	Responders  []OpsGenieMessageResponder  `json:"responders"`
	Entity      string 						`json:"entity"`
}

type OpsGenieMessageResponder struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

func main() {
	options := parseOptions()
	reader  := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString(utils.EOT_B)
		utils.CatchError(err)

		message    := utils.UnMarshalMessage(text)

		// create ops message
		opsMessage := new(OpsGenieMessage)
		opsMessage.Message = message.Message
		opsMessage.Entity  = message.HostName

		responder := OpsGenieMessageResponder{Id: options.ResponderId, Type: options.ResponderType}
		opsMessage.Responders = append(opsMessage.Responders, responder)

		b, err := json.Marshal(opsMessage)
		utils.CatchError(err)

		postData := string(b)

		req, err := http.NewRequest("POST", "https://api.eu.opsgenie.com/v2/alerts", bytes.NewBufferString(postData))
		utils.CatchError(err)

		req.Header.Add("Content-Type", "application/json")
		req.Header.Add("Authorization", "GenieKey " + options.ApiToken)

		client := http.Client{}
		_, err = client.Do(req)
		utils.CatchError(err)

		_, err = os.Stdout.WriteString(text)
		utils.CatchError(err)
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	strPtr1 := utils.GetConfigString("apiToken", "", "Api token", "opsgenie")
	strPtr2 := utils.GetConfigString("apiEndpoint", "", "Api endpoint (example: https://...)", "opsgenie")
	strPtr3 := utils.GetConfigString("responderType", "", "Responder type", "opsgenie")
	strPtr4 := utils.GetConfigString("responderId", "", "Responder id", "opsgenie")

	flag.Parse()

	options.ApiToken      = *strPtr1
	options.ApiEndpoint   = *strPtr2
	options.ResponderType = *strPtr3
	options.ResponderId   = *strPtr4

	if len(options.ApiToken) == 0 {
		panic("Api token required")
	}
	if len(options.ApiEndpoint) == 0 {
		panic("Api endpoint required")
	}
	if len(options.ResponderType) == 0 {
		panic("ResponderType required")
	}
	if len(options.ResponderId) == 0 {
		panic("ResponderId type required")
	}


	return options
}