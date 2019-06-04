package main

import (
	"climonitoring/utils"
	"flag"
	"net"
	"os"
	"time"
)

type CliOptions struct {
	UpdateInterval int
	Network        string
	Address        string
}

func main() {
	var result string
	options := parseOptions()

	if len(options.Address) == 0 {
		panic("Address required")
	}

	for {
		c, err := net.Dial(options.Network, options.Address)

		if err != nil {
			result = "false"
		} else {
			err = c.Close()
			utils.CatchError(err)
			result = "true"
		}

		_, err = os.Stdout.WriteString(result + utils.EOT_S)
		utils.CatchError(err)

		time.Sleep(time.Duration(options.UpdateInterval) * time.Second)
	}
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr  := utils.GetConfigInt("i", 1, "Update interval", "tcp")
	strPtr  := utils.GetConfigString("a", "", "Address (host:port)", "tcp")
	strPtr2 := utils.GetConfigString("n", "tcp", "Network to use", "tcp")

	flag.Parse()

	options.UpdateInterval = *intPtr
	options.Address        = *strPtr
	options.Network        = *strPtr2

	return options
}