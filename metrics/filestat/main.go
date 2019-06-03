package main

import (
	"climonitoring/utils"
	"encoding/json"
	"flag"
	"os"
	"time"
)

type Result struct {
	Size int64
	ModTime int64
	HasChanged bool
}

type CliOptions struct {
	UpdateInterval int
	FilePath string
}

func main() {
	options := parseOptions()
	result  := new(Result)

	if len(options.FilePath) == 0 {
		panic("File path required")
	}

	for {
		// save previous result (copy)
		var oldResult = *result

		fileInfo, err := os.Stat(options.FilePath)

		if err != nil {
			panic(err)
		}

		fileInfo.Size()

		result.Size       = fileInfo.Size()
		result.ModTime    = fileInfo.ModTime().Unix()
		result.HasChanged = oldResult.ModTime > 0 && result.ModTime != oldResult.ModTime

		// make json
		b, err := json.Marshal(result)
		utils.CatchError(err)

		// write json
		_, err = os.Stdout.WriteString(string(b) + utils.EOT_S)
		utils.CatchError(err)

		time.Sleep(time.Duration(options.UpdateInterval) * time.Second)
	}

}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr := utils.GetConfigInt("i", 1, "Update interval", "filestat")
	strPtr := utils.GetConfigString("f", "", "File path", "filestat")


	flag.Parse()

	options.UpdateInterval = *intPtr
	options.FilePath       = *strPtr

	return options
}