package main

import (
	"climonitoring/utils"
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Result struct {
	TotalRam   uint64
	FreeRam    uint64
	AvailRam   uint64
	BuffersRam uint64
	CacheRam   uint64
	UsedRam    uint64
}

type CliOptions struct {
	UpdateInterval int
}

func main() {
	options := parseOptions()
	result  := new(Result)

	for {
		b, err := ioutil.ReadFile("/proc/meminfo")
		utils.CatchError(err)

		content := string(b)

		r, _  := regexp.Compile(`MemFree:\s+(\d+)`)
		match := r.FindStringSubmatch(content)
		intVal, _ := strconv.Atoi(match[1])
		result.FreeRam = uint64(intVal * 1024)

		r, _  = regexp.Compile(`MemTotal:\s+(\d+)`)
		match = r.FindStringSubmatch(content)
		intVal, _ = strconv.Atoi(match[1])
		result.TotalRam = uint64(intVal * 1024)

		r, _  = regexp.Compile(`MemAvailable:\s+(\d+)`)
		match = r.FindStringSubmatch(content)
		intVal, _ = strconv.Atoi(match[1])
		result.AvailRam = uint64(intVal * 1024)

		r, _  = regexp.Compile(`Buffers:\s+(\d+)`)
		match = r.FindStringSubmatch(content)
		intVal, _ = strconv.Atoi(match[1])
		result.BuffersRam = uint64(intVal * 1024)

		r, _  = regexp.Compile(`Cached:\s+(\d+)`)
		match = r.FindStringSubmatch(content)
		intVal, _ = strconv.Atoi(match[1])
		result.CacheRam = uint64(intVal * 1024)

		result.UsedRam = result.TotalRam - result.FreeRam - result.BuffersRam - result.CacheRam

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