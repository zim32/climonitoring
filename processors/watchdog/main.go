package main

import (
	"bufio"
	"climonitoring/utils"
	"flag"
	"os"
	"time"
)

type CliOptions struct {
	MaxInterval int64
}

func main() {
	options  := parseOptions()
	reader   := bufio.NewReader(os.Stdin)
	lastTime := time.Now().Unix()

	go func() {
		for {
			if time.Now().Unix() - lastTime > options.MaxInterval {
				_, err := os.Stdout.WriteString("true" + utils.EOT_S)
				utils.CatchError(err)
			}
			time.Sleep(time.Duration(1 * time.Second))
		}
	}()

	for {
		_, err := utils.GetNewLine(reader)
		utils.CatchError(err)
		lastTime = time.Now().Unix()
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr := utils.GetConfigInt("i", 10, "Max no action interval (seconds)", "watchdog")

	flag.Parse()

	options.MaxInterval = int64(*intPtr)


	return options
}