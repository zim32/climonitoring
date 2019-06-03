package main

import (
	"climonitoring/utils"
	"bufio"
	"flag"
	"os"
	"regexp"
	"strings"
)

type CliOptions struct {
	Pattern string
}


func main() {
	options := parseOptions()
	reader  := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString(utils.EOT_B)
		utils.CatchError(err)

		val      := strings.Trim(text, utils.EOT_S)
		regex, _ := regexp.Compile(options.Pattern)

		if !regex.MatchString(val) {
			continue
		}

		_, err = os.Stdout.WriteString(text)
		utils.CatchError(err)
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	strPtr := utils.GetConfigString("e", ".*", "Pattern", "regex")

	flag.Parse()

	options.Pattern = *strPtr

	return options
}