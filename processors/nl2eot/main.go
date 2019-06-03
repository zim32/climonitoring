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
		text, err := reader.ReadString('\n')
		utils.CatchEof(err, text)

		text = strings.Replace(text,"\n", utils.EOT_S, -1)

		_, err = os.Stdout.WriteString(text)

		if err != nil {
			log.Print(err)
			os.Exit(0)
		}
	}
}