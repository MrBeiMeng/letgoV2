package logging

import (
	"fmt"
	"letgoV2/system_code/pkg/setting"
	"log"
	"os"
	"path"
	"time"
)

var (
	LogSavePath = "logs/"
	LogSaveName = "log"
	LogFileExt  = "log"
	TimeFormat  = "20060102"
)

func getLogFilePath() string {

	logDir := setting.LogDir
	if setting.NotRootStart {
		logDir = "."
	}

	fullLogSavePath := path.Join(logDir, LogSavePath)

	return fmt.Sprintf("%s/", fullLogSavePath)
}

func GetLogFileFullPath(prefixPath string) string {
	suffixPath := fmt.Sprintf("%s%s.%s", LogSaveName, time.Now().Format(TimeFormat), LogFileExt)

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func OpenLogFile(filePath string) *os.File {

	_, err := os.Stat(filePath)
	switch {
	case os.IsNotExist(err):
		mkDir()
	case os.IsPermission(err):
		log.Fatalf("Permission :%v", err)
	}

	handle, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Fail to OpenFile :%v", err)
	}

	return handle
}

func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir+"/"+getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}
