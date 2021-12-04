package main

import (
	"encoding/json"
	"flag"
	"os"
	"syscall"
	"time"

	"github.com/zim32/climonitoring/utils"
)

type Result struct {
	LoadAvg1  float32
	LoadAvg5  float32
	LoadAvg15 float32
}

type CliOptions struct {
	UpdateInterval int
}

const LoadAvgShift = 65536

func main() {
	options := parseOptions()
	result := new(Result)
	sysInfo := new(syscall.Sysinfo_t)

	for {
		// make syscall
		err := syscall.Sysinfo(sysInfo)

		if err != nil {
			panic(err)
		}

		result.LoadAvg1 = float32(sysInfo.Loads[0]) / LoadAvgShift
		result.LoadAvg5 = float32(sysInfo.Loads[1]) / LoadAvgShift
		result.LoadAvg15 = float32(sysInfo.Loads[2]) / LoadAvgShift

		// make json
		b, err := json.Marshal(result)

		// write json
		_, err = os.Stdout.WriteString(string(b) + utils.EOT_S)
		utils.CatchError(err)

		time.Sleep(time.Duration(options.UpdateInterval) * time.Second)
	}

}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr := utils.GetConfigInt("i", 1, "Update interval", "cpu")

	flag.Parse()

	options.UpdateInterval = *intPtr

	return options
}
