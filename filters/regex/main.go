package main

import (
	"bufio"
	"flag"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/zim32/climonitoring/utils"
)

type CliOptions struct {
	Pattern     string
	OutTemplate string
	Invert      bool
}

func main() {
	options := parseOptions()
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		val := strings.Trim(text, utils.EOT_S)
		regex := regexp.MustCompile(options.Pattern)

		if len(options.OutTemplate) > 0 {
			matches := regex.FindStringSubmatch(val)

			if len(matches) == 0 && !options.Invert {
				continue
			}

			if len(matches) > 0 && options.Invert {
				continue
			}

			if options.Invert {
				_, err = os.Stdout.WriteString(options.OutTemplate + utils.EOT_S)
				utils.CatchError(err)
				continue
			}

			outString := options.OutTemplate

			for i, val := range matches {
				outString = strings.Replace(outString, "{"+strconv.Itoa(i)+"}", val, -1)
			}

			_, err = os.Stdout.WriteString(outString + utils.EOT_S)
			utils.CatchError(err)
		} else {
			ok := regex.MatchString(val)
			if !ok && !options.Invert {
				continue
			}
			if ok && options.Invert {
				continue
			}

			_, err = os.Stdout.WriteString(text)
			utils.CatchError(err)
		}
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	strPtr1 := utils.GetConfigString("e", ".*", "Pattern", "regex")
	strPtr2 := utils.GetConfigString("o", "", "Output template", "regex")
	booPtr1 := utils.GetConfigBool("invert", false, "Invert", "regex")

	flag.Parse()

	options.Pattern = *strPtr1
	options.OutTemplate = *strPtr2
	options.Invert = *booPtr1

	return options
}
