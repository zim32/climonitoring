package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/smtp"
	"os"

	"github.com/zim32/climonitoring/utils"
)

type CliOptions struct {
	SmtpHost     string
	SmtpPort     string
	UserName     string
	UserPassword string
	SendFrom     string
	SendTo       string
}

func main() {
	options := parseOptions()
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		message := utils.UnMarshalMessage(text)

		bodyText := `
Severity: %s
Message: %s
Host: %s
Created: %s
`
		bodyText = fmt.Sprintf(bodyText, message.Severity, message.Message, message.HostName, message.Created)
		subjectText := fmt.Sprintf("[%s] New alert from %s", message.Severity, message.HostName)
		msgHeader := fmt.Sprintf("To: %s\r\nSubject: %s\r\n", options.SendTo, subjectText)
		msg := msgHeader + "\r\n" + bodyText + "\r\n"

		auth := smtp.PlainAuth(
			"",
			options.UserName,
			options.UserPassword,
			options.SmtpHost,
		)

		err = smtp.SendMail(options.SmtpHost+":"+options.SmtpPort, auth, options.SendFrom, []string{options.SendTo}, []byte(msg))
		utils.CatchError(err)

		_, err = os.Stdout.WriteString(text)
		utils.CatchError(err)
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	strPtr1 := utils.GetConfigString("host", "", "SMTP host", "smtp")
	strPtr2 := utils.GetConfigString("port", "", "SMTP port", "smtp")
	strPtr3 := utils.GetConfigString("userName", "", "SMTP user name", "smtp")
	strPtr4 := utils.GetConfigString("userPass", "", "SMTP user password", "smtp")
	strPtr5 := utils.GetConfigString("from", "", "Send from", "smtp")
	strPtr6 := utils.GetConfigString("to", "", "Send to", "smtp")

	flag.Parse()

	options.SmtpHost = *strPtr1
	options.SmtpPort = *strPtr2
	options.UserName = *strPtr3
	options.UserPassword = *strPtr4
	options.SendFrom = *strPtr5
	options.SendTo = *strPtr6

	utils.ValidateEmptyString(options.SmtpHost, "SMTP host required")
	utils.ValidateEmptyString(options.SmtpPort, "SMTP port required")
	utils.ValidateEmptyString(options.UserName, "SMTP user name required")
	utils.ValidateEmptyString(options.UserPassword, "SMTP user password required")
	utils.ValidateEmptyString(options.SendFrom, "Send from required")
	utils.ValidateEmptyString(options.SendFrom, "Send to required")

	return options
}
