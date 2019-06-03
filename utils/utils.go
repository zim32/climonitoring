package utils

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-ini/ini"
	"io"
	"log"
	"os"
	"strings"
	"time"
)

const EOT_B = 4
const EOT_S = "\x04"

type Message struct {
	Severity string
	Message  string
	Created  time.Time
	HostName string
}

func DumpStruct(s interface{}) {
	fmt.Printf("%+v\n", s)
}

func UnMarshalMessage(text string) Message {
	text = strings.Trim(text, EOT_S)
	var message Message

	err := json.Unmarshal([]byte(text), &message)
	CatchError(err)

	return message
}

func CatchError(err error) {
	if err == nil {
		return
	}

	if err == io.EOF {
		log.Print("Use CatchEof for stdin processing\n")
		os.Exit(0)
	} else {
		panic(err)
	}
}

func CatchEof(err error, text string) {
	if err == nil {
		return
	}

	if err == io.EOF && len(text) == 0 {
		os.Exit(0)
	}
}

func FileExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil { return true, nil }
	if os.IsNotExist(err) { return false, nil }
	return true, err
}

func getConfigFilePath() (string, bool)  {
	var iniFile string

	if ok, _ := FileExists("cm.config.ini"); ok {
		iniFile = "cm.config.ini"
	} else if ok, _ := FileExists("/etc/cm/config.ini"); ok {
		iniFile = "/etc/cm/config.ini"
	}

	if len(iniFile) == 0 {
		return "", false
	}

	return iniFile, true
}

func GetIniString(section string, key string, def string) string  {
	if iniFile, ok := getConfigFilePath(); ok {
		cfg, err := ini.Load(iniFile)
		if err != nil {
			return def
		}

		return cfg.Section(section).Key(key).MustString(def)
	} else {
		return def
	}
}

func GetIniInt(section string, key string, def int) int  {
	if iniFile, ok := getConfigFilePath(); ok {
		cfg, err := ini.Load(iniFile)
		if err != nil {
			return def
		}

		return cfg.Section(section).Key(key).MustInt(def)
	} else {
		return def
	}
}

func GetConfigString(name string, def string, usage string, section string) *string  {
	defVal := GetIniString(section, name, def)
	ptr    := flag.String(name, defVal, usage)

	return ptr
}


func GetConfigInt(name string, def int, usage string, section string) *int  {
	defVal := GetIniInt(section, name, def)
	ptr    := flag.Int(name, defVal, usage)

	return ptr
}

func NewMessage() *Message {
	msg := new(Message)

	osHostName, err := os.Hostname()
	CatchError(err)

	msg.HostName = GetIniString("global", "host_name", osHostName)

	return msg
}

func GetNewLine(reader *bufio.Reader) (string, error) {
	text, err := reader.ReadString(EOT_B)
	CatchEof(err, text)

	if err == io.EOF {
		text = strings.Trim(text, "\n") + EOT_S
	}

	return text, nil
}