package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/zim32/climonitoring/utils"
)

type CliOptions struct {
	UpdateInterval int
	Pid            int
	Name           string
}

type ResultItem struct {
	Pid             string
	ProcessName     string
	CommandLine     string
	ProcessState    string
	NumberOfThreads int
	MemRss          int64
	MemRssAnon      int64
	MemRssFile      int64
	MemRssShared    int64
	MemRssOwn       int64
	MemVirtual      int64
	CoreDumping     int
	NetInBytes      uint64
	NetOutBytes     uint64
}

func main() {
	options := parseOptions()

	for {
		var pids []int

		if options.Pid > 0 {
			pids = append(pids, options.Pid)
		} else {
			pids = getPidsByPattern(options.Name)
		}

		result := new(ResultItem)

		for _, pid := range pids {
			resultItem := new(ResultItem)

			parseProcStatus(pid, resultItem)
			parseProcNetstat(pid, resultItem)

			if len(pids) == 1 {
				parseCommandLine(pid, resultItem)
				result.ProcessState = resultItem.ProcessState
			}

			resultItem.Pid = strconv.Itoa(pid)

			result.Pid += resultItem.Pid + "|"
			result.CommandLine += resultItem.CommandLine + "|"
			result.MemRssShared += resultItem.MemRssShared
			result.MemRss += resultItem.MemRss
			result.MemRssOwn += resultItem.MemRssOwn
			result.MemRssFile += resultItem.MemRssFile
			result.MemRssAnon += resultItem.MemRssAnon
			result.MemVirtual += resultItem.MemVirtual
			result.NumberOfThreads += resultItem.NumberOfThreads
			result.CoreDumping += resultItem.CoreDumping
			result.CommandLine += resultItem.CommandLine
			result.NetInBytes += resultItem.NetInBytes
			result.NetOutBytes += resultItem.NetOutBytes
		}

		result.Pid = strings.Trim(result.Pid, "|")
		result.CommandLine = strings.Trim(result.CommandLine, "|")

		// make json
		b, err := json.Marshal(result)

		// write json
		_, err = os.Stdout.WriteString(string(b) + utils.EOT_S)
		utils.CatchError(err)

		time.Sleep(time.Duration(options.UpdateInterval) * time.Second)
	}
}

func getPidsByPattern(pattern string) []int {
	procFiles, err := ioutil.ReadDir("/proc")
	utils.CatchError(err)
	var pids []int

	for _, file := range procFiles {
		if !file.IsDir() {
			continue
		}
		if ok, _ := regexp.MatchString(`\d+`, file.Name()); !ok {
			continue
		}

		processPid, err := strconv.Atoi(file.Name())
		utils.CatchError(err)

		b, err := ioutil.ReadFile("/proc/" + file.Name() + "/comm")
		utils.CatchError(err)

		procName := string(b)

		if ok, _ := regexp.MatchString(pattern, procName); ok {
			pids = append(pids, processPid)
		}
	}

	return pids
}

func parseProcStatus(pid int, resultItem *ResultItem) {
	b, err := ioutil.ReadFile("/proc/" + strconv.Itoa(pid) + "/status")
	utils.CatchError(err)

	content := string(b)

	for _, line := range strings.Split(content, "\n") {
		r := regexp.MustCompile(`State:\s+(.+)`)
		if match := r.FindStringSubmatch(line); len(match) > 0 {
			resultItem.ProcessState = match[1]
		}

		r = regexp.MustCompile(`VmSize:\s+(\d+)`)
		if match := r.FindStringSubmatch(line); len(match) > 0 {
			if i, err := strconv.Atoi(match[1]); err == nil {
				resultItem.MemVirtual = int64(i) * 1024
			}
		}

		r = regexp.MustCompile(`VmRSS:\s+(\d+)`)
		if match := r.FindStringSubmatch(line); len(match) > 0 {
			if i, err := strconv.Atoi(match[1]); err == nil {
				resultItem.MemRss = int64(i) * 1024
			}
		}

		r = regexp.MustCompile(`RssAnon:\s+(\d+)`)
		if match := r.FindStringSubmatch(line); len(match) > 0 {
			if i, err := strconv.Atoi(match[1]); err == nil {
				resultItem.MemRssAnon = int64(i) * 1024
			}
		}

		r = regexp.MustCompile(`RssFile:\s+(\d+)`)
		if match := r.FindStringSubmatch(line); len(match) > 0 {
			if i, err := strconv.Atoi(match[1]); err == nil {
				resultItem.MemRssFile = int64(i) * 1024
			}
		}

		r = regexp.MustCompile(`RssShmem:\s+(\d+)`)
		if match := r.FindStringSubmatch(line); len(match) > 0 {
			if i, err := strconv.Atoi(match[1]); err == nil {
				resultItem.MemRssShared = int64(i) * 1024
			}
		}

		r = regexp.MustCompile(`CoreDumping:\s+(\d+)`)
		if match := r.FindStringSubmatch(line); len(match) > 0 {
			if i, err := strconv.Atoi(match[1]); err == nil {
				resultItem.CoreDumping = i
			}
		}

		r = regexp.MustCompile(`Threads:\s+(\d+)`)
		if match := r.FindStringSubmatch(line); len(match) > 0 {
			if i, err := strconv.Atoi(match[1]); err == nil {
				resultItem.NumberOfThreads = i
			}
		}
	}

	resultItem.MemRssOwn = resultItem.MemRss - resultItem.MemRssShared
}

func parseCommandLine(pid int, resultItem *ResultItem) {
	b, err := ioutil.ReadFile("/proc/" + strconv.Itoa(pid) + "/cmdline")
	utils.CatchError(err)

	content := string(b)
	resultItem.CommandLine = content
}

func parseProcNetstat(pid int, resultItem *ResultItem) {
	b, err := ioutil.ReadFile("/proc/" + strconv.Itoa(pid) + "/net/netstat")
	utils.CatchError(err)

	content := string(b)

	r := regexp.MustCompile(`IpExt: \d+.*`)
	match := r.FindString(content)

	tmp := strings.Split(match, " ")

	i, err := strconv.Atoi(tmp[7])
	utils.CatchError(err)
	resultItem.NetInBytes = uint64(i)

	i, err = strconv.Atoi(tmp[7])
	utils.CatchError(err)
	resultItem.NetOutBytes = uint64(i)
}

func parseOptions() *CliOptions {
	options := new(CliOptions)

	intPtr1 := flag.Int("i", 1, "Update interval")
	intPtr2 := utils.GetConfigInt("pid", -1, "Process PID", "procinfo")
	strPtr := utils.GetConfigString("name", "", "Process name (regex pattern)", "procinfo")

	flag.Parse()

	options.UpdateInterval = *intPtr1
	options.Pid = *intPtr2
	options.Name = *strPtr

	if options.Pid == -1 && len(options.Name) == 0 {
		panic("-pid or -name parameter required")
	}

	return options
}
