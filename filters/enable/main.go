package main

import (
	"bufio"
	"climonitoring/utils"
	"flag"
	"github.com/fsnotify/fsnotify"
	"io/ioutil"
	"os"
	"strings"
)

type CliOptions struct {
	FilePath string
	TrueVal  string
}

var fileContent string

func main() {
	options := parseOptions()
	reader  := bufio.NewReader(os.Stdin)

	readFileContent(options.FilePath)

	watcher, err := fsnotify.NewWatcher()
	utils.CatchError(err)

	go func() {
		// watch for content file changes
		for {
			select {
				case event, ok := <-watcher.Events:
					if !ok {
						return
					}
					if event.Op & fsnotify.Write == fsnotify.Write {
						readFileContent(options.FilePath)
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return
					}
					utils.CatchError(err)
			}
		}
	}()

	err = watcher.Add(options.FilePath)
	utils.CatchError(err)


	for {
		text, err := utils.GetNewLine(reader)

		if fileContent != options.TrueVal {
			continue
		}

		_, err = os.Stdout.WriteString(text)
		utils.CatchError(err)
	}
}

func readFileContent(path string) {
	b, err := ioutil.ReadFile(path)
	utils.CatchError(err)

	fileContent = strings.Trim(string(b), "\n")
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	strPtr1 := utils.GetConfigString("f", "", "File path to read content from", "enable")
	strPtr2 := utils.GetConfigString("s", "1", "String interpret as true value", "enable")

	flag.Parse()

	options.FilePath = *strPtr1
	options.TrueVal  = *strPtr2

	utils.ValidateEmptyString(options.FilePath, "File path required")

	return options
}