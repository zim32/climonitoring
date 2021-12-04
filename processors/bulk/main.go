package main

import (
	"bufio"
	"bytes"
	"flag"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/zim32/climonitoring/utils"
)

type CliOptions struct {
	BufferSize    int
	FlushInterval int
}

func main() {
	options := parseOptions()
	buffer := make([]string, 0)
	reader := bufio.NewReader(os.Stdin)
	mutex := new(sync.Mutex)

	var flushTime int64 = 0

	if options.FlushInterval > 0 {
		// flush loop
		go func() {
			for {
				mutex.Lock()

				if len(buffer) == 0 || flushTime == 0 {
					mutex.Unlock()
					time.Sleep(time.Duration(1 * time.Second))
					continue
				}

				if (time.Now().Unix() - flushTime) > int64(options.FlushInterval) {
					flushBuffer(&buffer)
					flushTime = 0
				}

				mutex.Unlock()

				time.Sleep(time.Duration(1 * time.Second))
			}
		}()
	}

	for {
		text, _ := utils.GetNewLine(reader)

		mutex.Lock()

		buffer = append(buffer, strings.Trim(text, utils.EOT_S))
		flushTime = time.Now().Unix()

		if len(buffer) >= options.BufferSize {
			flushBuffer(&buffer)
		}

		mutex.Unlock()
	}
}

func flushBuffer(buffer *[]string) {
	result := new(bytes.Buffer)

	for _, line := range *buffer {
		result.WriteString(line + "\n")
	}

	_, err := os.Stdout.WriteString(result.String() + utils.EOT_S)
	utils.CatchError(err)

	*buffer = make([]string, 0)
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr := utils.GetConfigInt("s", 10, "Buffer size", "bulk")
	intPtr2 := utils.GetConfigInt("fi", 0, "Max flush interval (seconds)", "bulk")

	flag.Parse()

	options.BufferSize = *intPtr
	options.FlushInterval = *intPtr2

	return options
}
