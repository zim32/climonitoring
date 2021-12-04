package main

import (
	"bufio"
	"flag"
	"os"

	"github.com/zim32/climonitoring/utils"
)

type CliOptions struct {
	SomeOption int64
}

func main() {
	//options := parseOptions()
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		_, err = os.Stdout.WriteString(text + utils.EOT_S)
		utils.CatchError(err)
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr := utils.GetConfigInt("cid", -1, "Chat ID", "telegram")

	flag.Parse()

	options.SomeOption = int64(*intPtr)

	return options
}
