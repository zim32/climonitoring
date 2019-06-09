package main

import (
	"climonitoring/utils"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Result struct {
	NetInBytes      uint64
	NetOutBytes     uint64
}

type CliOptions struct {
	UpdateInterval int
}

func main() {
	options := parseOptions()
	result  := new(Result)

	for {
		b, err := ioutil.ReadFile("/proc/net/netstat")
		utils.CatchError(err)

		content := string(b)

		r := regexp.MustCompile(`IpExt: \d+.*`)
		match := r.FindString(content)

		tmp := strings.Split(match, " ")

		i, err := strconv.Atoi(tmp[7])
		utils.CatchError(err)
		result.NetInBytes = uint64(i)

		i, err = strconv.Atoi(tmp[7])
		utils.CatchError(err)
		result.NetOutBytes = uint64(i)

		// make json
		b, err = json.Marshal(result)

		// write json
		_, err = os.Stdout.WriteString(string(b) + utils.EOT_S)
		utils.CatchError(err)

		time.Sleep(time.Duration(options.UpdateInterval) * time.Second)
	}

}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr := flag.Int("i", 1, "Update interval")

	flag.Parse()

	options.UpdateInterval = *intPtr

	return options
}