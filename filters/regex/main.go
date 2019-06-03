package main

import (
	"bufio"
	"climonitoring/utils"
	"flag"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type CliOptions struct {
	Pattern     string
	OutTemplate string
}


func main() {
	options := parseOptions()
	reader  := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		val   := strings.Trim(text, utils.EOT_S)
		regex := regexp.MustCompile(options.Pattern)

		if len(options.OutTemplate) > 0 {
			matches := regex.FindStringSubmatch(val)

			if len(matches) == 0 {
				continue
			}

			outString := options.OutTemplate

			for i, val := range matches {
				outString = strings.Replace(outString, "{" + strconv.Itoa(i) + "}", val, -1)
			}

			_, err = os.Stdout.WriteString(outString + utils.EOT_S)
			utils.CatchError(err)
		} else {
			if !regex.MatchString(val) {
				continue
			}

			_, err = os.Stdout.WriteString(text)
			utils.CatchError(err)
		}
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	strPtr1 := utils.GetConfigString("e", ".*", "Pattern","regex")
	strPtr2 := utils.GetConfigString("o", "", "Output template", "regex")

	flag.Parse()

	options.Pattern     = *strPtr1
	options.OutTemplate = *strPtr2

	return options
}