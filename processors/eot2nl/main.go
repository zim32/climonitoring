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
		text, err := reader.ReadString(utils.EOT_B)

		if err != nil {
			log.Print(err)
			os.Exit(0)
		}

		text = strings.Replace(text, utils.EOT_S, "\n", -1)

		_, err = os.Stdout.WriteString(text)

		if err != nil {
			log.Print(err)
			os.Exit(0)
		}
	}
}