package clog

import (
	"encoding/json"
	"fmt"
	"github.com/chu108/tool"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Level int

type LogMsg struct {
	Prifix string                 `json:"prifix"`
	Level  string                 `json:"level"`
	File   string                 `json:"file"`
	Line   int                    `json:"line"`
	Time   string                 `json:"time"`
	Param  map[string]interface{} `json:"param"`
}

var (
	logger             *log.Logger
	F                  *os.File
	DefaultPrefix      = ""
	DefaultCallerDepth = 2
	logPrefix          = "-"
	levelFlags         = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := getLogPath()
	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		tool.Err(err)
	}
	logger = log.New(handle, DefaultPrefix, 0)
}

func Debug(v ...interface{}) {
	printJson(DEBUG, v...)
}

func Info(v ...interface{}) {
	printJson(INFO, v...)
}

func Warn(v ...interface{}) {
	printJson(WARNING, v...)
}

func Error(v ...interface{}) {
	printJson(ERROR, v...)
}

func Fatal(v ...interface{}) {
	printJson(FATAL, v...)
}

func SetPrefix(prefix string) {
	logPrefix = prefix
}

func printJson(level Level, kv ...interface{}) {
	_, file, line, _ := runtime.Caller(DefaultCallerDepth)
	lg := &LogMsg{
		Prifix: logPrefix,
		Time:   time.Now().Format("2006-01-02 15:04:05"),
		Level:  levelFlags[level],
		File:   filepath.Base(file),
		Line:   line,
		Param:  make(map[string]interface{}),
	}
	if len(kv)%2 != 0 {
		kv = append(kv, "-")
	}
	for i := 0; i < len(kv); i += 2 {
		lg.Param[fmt.Sprintf("%v", kv[i])] = kv[i+1]
	}
	msgByte, _ := json.Marshal(lg)
	msgStr := tool.BytesToStr(msgByte)
	switch level {
	case FATAL:
		tool.Err(msgStr)
		logger.Fatalln(msgStr)
	case ERROR:
		tool.Err(msgStr)
		logger.Println(msgStr)
	case DEBUG, INFO, WARNING:
		tool.Info(msgStr)
		logger.Println(msgStr)
	}
}

func getLogPath() string {
	dir, err := os.Getwd()
	if err != nil {
		dir = "."
	}
	filePath := fmt.Sprintf("%s/log/%s%s.%s", dir, logPrefix, time.Now().Format("20060102"), "log")
	err = tool.CreateFileByNot(dir)
	if err != nil {
		tool.Err(err)
	}
	return filePath
}
