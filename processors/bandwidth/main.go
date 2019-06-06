package main

import (
	"bufio"
	"climonitoring/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader  := bufio.NewReader(os.Stdin)
	start   := time.Now()

	var prevVal float64

	for {
		text, err := utils.GetNewLine(reader)
		now       := time.Now()

		val, err := strconv.ParseFloat(strings.Trim(text, utils.EOT_S), 64)
		utils.CatchError(err)

		if prevVal == 0 {
			prevVal = val
			start   = time.Now()
			continue
		}

		if val - prevVal < 0 {
			continue
		}

		bandwidth :=  int64((val - prevVal) / now.Sub(start).Seconds())
		start      = time.Now()
		prevVal    = val

		_, err = os.Stdout.WriteString(strconv.FormatInt(bandwidth, 10) + utils.EOT_S)
		utils.CatchError(err)
	}
}