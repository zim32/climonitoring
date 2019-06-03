package main

import (
	"climonitoring/utils"
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		val := strings.ToLower(text)
		val = strings.Trim(text, utils.EOT_S)

		if val == "true" || val == "1" {
			_, err = os.Stdout.WriteString(text)

			if err != nil {
				log.Print(err)
				os.Exit(0)
			}
		}
	}
}