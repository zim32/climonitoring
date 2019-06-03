package main

import (
	"climonitoring/utils"
	"bufio"
	"flag"
	"os"
	"time"
)

type CliOptions struct {
	Interval int64
}

func main() {
	options := parseOptions()
	reader  := bufio.NewReader(os.Stdin)
	var lastTime int64 = 0

	for {
		text, err := utils.GetNewLine(reader)

		if time.Now().Unix() - lastTime < options.Interval {
			continue
		}

		_, err = os.Stdout.WriteString(text)
		utils.CatchError(err)

		lastTime = time.Now().Unix()
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr  := utils.GetConfigInt("i", 10, "Debounce interval (seconds)", "debounce")

	flag.Parse()

	options.Interval = int64(*intPtr)

	return options
}