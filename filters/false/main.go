package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/zim32/climonitoring/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		val := strings.ToLower(text)
		val = strings.Trim(text, utils.EOT_S)

		if val == "false" || val == "0" {
			_, err = os.Stdout.WriteString(text)

			if err != nil {
				log.Print(err)
				os.Exit(0)
			}
		}
	}
}
