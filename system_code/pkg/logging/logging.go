package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Level int

var (
	F *os.File

	DefaultPrefix      = ""
	DefaultCallerDepth = 2

	logger     *log.Logger
	logPrefix  = ""
	levelFlags = []string{"DEBUG", "INFO", "WARN", "ERROR", "FATAL"}
)

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
	FATAL
)

func init() {
	filePath := GetLogFileFullPath(getLogFilePath())
	F = OpenLogFile(filePath)

	multiWriter := io.MultiWriter(os.Stdout, F)

	logger = log.New(multiWriter, DefaultPrefix, log.Ltime)
}

func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(replaceNextLine(v))
}

func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(replaceNextLine(v...))
}

func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(replaceNextLine(v))
}

func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(replaceNextLine(v))
}

func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Fatalln(v)
}

func replaceNextLine(v ...interface{}) string {
	str := fmt.Sprint(v...)
	return strings.ReplaceAll(str, "\n", "\\n")
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if !ok {
		logPrefix = fmt.Sprintf("  %s  |  %s:%d  |  ", levelFlags[level], filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf(" %s | %s | ", levelFlags[level], filepath.Base(filepath.Dir(file)))
	}

	logger.SetPrefix(logPrefix)
}
