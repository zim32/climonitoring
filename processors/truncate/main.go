package main

import (
	"climonitoring/utils"
	"bufio"
	"flag"
	"os"
	"strings"
)

type CliOptions struct {
	MaxLength int
}

func main() {
	options := parseOptions()
	reader  := bufio.NewReader(os.Stdin)

	for {
		text, err := reader.ReadString(utils.EOT_B)
		utils.CatchError(err)

		text = strings.Trim(text, utils.EOT_S)

		if options.MaxLength > 0 && len(text) > options.MaxLength {
			text = text[0:options.MaxLength] + "...Truncated"
		}

		_, err = os.Stdout.WriteString(text + utils.EOT_S)
		utils.CatchError(err)
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr  := utils.GetConfigInt("l", 0, "Max length", "truncate")

	flag.Parse()

	options.MaxLength = *intPtr

	return options
}