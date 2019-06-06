package main

import (
	"bufio"
	"climonitoring/utils"
	"flag"
	"os"
	"strconv"
	"strings"
	"time"
)

type CliOptions struct {
	TimeInterval uint
}

func main() {
	options := parseOptions()
	reader  := bufio.NewReader(os.Stdin)
	start   := time.Now()

	var buffer []float64

	for {
		text, err := utils.GetNewLine(reader)
		now       := time.Now()

		val, err := strconv.ParseFloat(strings.Trim(text, utils.EOT_S), 64)
		utils.CatchError(err)

		buffer = append(buffer, val)

		if uint(now.Unix() - start.Unix()) > options.TimeInterval {
			avg := int64(calculateAverage(&buffer))

			start  = now
			buffer = nil

			_, err = os.Stdout.WriteString(strconv.FormatInt(avg, 10) + utils.EOT_S)
			utils.CatchError(err)
		}
	}
}

func calculateAverage(buffer *[]float64) float64 {
	var sum float64

	for _, item := range *buffer {
		sum += item
	}

	avg := sum / float64(len(*buffer))

	return avg
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr := utils.GetConfigInt("i", 10, "Time interval", "average")

	flag.Parse()

	options.TimeInterval = uint(*intPtr)

	return options
}