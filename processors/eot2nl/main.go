package main

import (
	"bufio"
	"climonitoring/utils"
	"log"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		text, err := utils.GetNewLine(reader)

		text = strings.Replace(text, utils.EOT_S, "\n", -1)

		_, err = os.Stdout.WriteString(text)

		if err != nil {
			log.Print(err)
			os.Exit(0)
		}
	}
}