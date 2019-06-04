package main

import (
	"bufio"
	"climonitoring/utils"
	"flag"
	"net/smtp"
	"os"
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
	reader  := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		auth := smtp.PlainAuth(
			"",
			options.UserName,
			options.UserPassword,
			options.SmtpHost,
		)

		err = smtp.SendMail(options.SmtpHost + ":" + options.SmtpPort, auth,options.SendFrom, []string{options.SendTo}, []byte(text))
		utils.CatchError(err)

		_, err = os.Stdout.WriteString(text + utils.EOT_S)
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

	options.SmtpHost     = *strPtr1
	options.SmtpPort     = *strPtr2
	options.UserName     = *strPtr3
	options.UserPassword = *strPtr4
	options.SendFrom     = *strPtr5
	options.SendTo       = *strPtr6

	utils.ValidateEmptyString(options.SmtpHost, "SMTP host required")
	utils.ValidateEmptyString(options.SmtpPort, "SMTP port required")
	utils.ValidateEmptyString(options.UserName, "SMTP user name required")
	utils.ValidateEmptyString(options.UserPassword, "SMTP user password required")
	utils.ValidateEmptyString(options.SendFrom, "Send from required")
	utils.ValidateEmptyString(options.SendFrom, "Send to required")

	return options
}