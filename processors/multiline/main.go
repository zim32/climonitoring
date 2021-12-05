package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/zim32/climonitoring/utils"
)

func main() {
	pattern := utils.GetConfigString("p", "", "Pattern to define new EOT split", "multiline")
	flag.Parse()

	if len(*pattern) == 0 {
		panic("Pattern required for multiline processor")
	}

	reader := bufio.NewReader(os.Stdin)

	buff := make([]byte, 100)
	regex := regexp.MustCompile(fmt.Sprintf("(%s)", *pattern))

	for {
		// read all available input into string
		var stringBuilder strings.Builder

		for {
			bytesRead, err := reader.Read(buff)

			if err != nil {
				log.Print(err)
				os.Exit(0)
			}

			slicedBuff := buff[0:bytesRead]
			stringBuilder.Write(slicedBuff)

			// prevent to much data to be buffered
			if stringBuilder.Len() > 1024*1024 {
				stringBuilder.Reset()
				stringBuilder.WriteString("Trying to buffer too many data")
				break
			}

			if bytesRead < len(buff) {
				// all data have been read
				break
			}
		}

		text := regex.ReplaceAllString(stringBuilder.String(), utils.EOT_S+"$1") + utils.EOT_S

		_, err := os.Stdout.WriteString(text)

		if err != nil {
			log.Print(err)
			os.Exit(0)
		}
	}
}
