package logging

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"letgoV2/system_code/pkg/e"
	"letgoV2/system_code/pkg/util/config_util"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

type Level int

var (
	logger *logrus.Logger
)

func init() {
	// 可能有多种不同的启动位置
	dir, _ := os.Getwd()

	logPath := filepath.Join(dir, "", "")

	logName := fmt.Sprintf("%s%s.%s", "log", time.Now().Format("20060102"), "log")
	if strings.Contains(dir, "your_code") {
		logPath = filepath.Join(dir, "./logs", logName)
	} else {
		logPath = filepath.Join(dir, config_util.Fields("Log").Get("Dir"), "logs", logName)
	}

	F, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}

	//multiWriter := io.MultiWriter(os.Stdout, F)

	// 设置logrus输出到标准输出
	//logrus.SetOutput(multiWriter)

	// 创建一个新的Logger实例
	logger = logrus.New()

	// 设置不同的formatter到不同的hook中
	consoleHook := NewColorfulConsoleHook(&logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "15:04:05",
	})
	//fileHook := NewFileHook(F, &logrus.JSONFormatter{})

	// 添加hook到Logger
	logger.AddHook(consoleHook)
	//logger.AddHook(fileHook)

	//logger.SetFormatter(&logrus.TextFormatter{
	//	ForceColors:   true,
	//	FullTimestamp: true,
	//})

	logger.SetOutput(F)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   "",
		DisableTimestamp:  false,
		DisableHTMLEscape: false,
		DataKey:           "",
		FieldMap:          nil,
		CallerPrettyfier:  nil,
		PrettyPrint:       false,
	})

}

// isTerminal 判断是否连接了TTY
func isTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func fieldsToMap(fields ...interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	for _, field := range fields {
		for key, value := range StructToMap(field) {
			result[key] = value
		}
	}

	return result
}

func Debug(msg interface{}, fields ...interface{}) {
	logger.WithFields(fieldsToMap(fields...)).Debugln(replaceNextLine(msg))
}

func Info(msg interface{}, fields ...interface{}) {
	logger.WithFields(fieldsToMap(fields...)).Infoln(replaceNextLine(msg))
}

func Warn(msg interface{}, fields ...interface{}) {
	logger.WithFields(fieldsToMap(fields...)).Warnln(replaceNextLine(msg))
}

func Error(msg interface{}, fields ...interface{}) {
	if customErr, ok := msg.(*e.LocationError); ok {
		msg = fmt.Sprintf("err occured: %s @ %s:%d", customErr.Error(), customErr.Location.FilePath, customErr.Location.Line)
	}

	logger.WithFields(fieldsToMap(fields...)).Errorln(replaceNextLine(msg))
}

func replaceNextLine(v ...interface{}) string {
	str := fmt.Sprint(v...)
	return strings.ReplaceAll(str, "\n", "\\n")
}

func StructToMap(obj interface{}) map[string]interface{} {
	// 获取结构体的反射值
	val := reflect.ValueOf(obj)

	// 如果传入的不是结构体对象，则返回空map
	if val.Kind() != reflect.Struct {
		return nil
	}

	// 创建一个map存储转换后的键值对
	result := make(map[string]interface{})

	// 获取结构体的类型
	typ := reflect.TypeOf(obj)

	// 遍历结构体的字段并添加到map中
	for i := 0; i < val.NumField(); i++ {
		// 获取字段的值
		fieldValue := val.Field(i).Interface()

		// 获取字段的名称，并转换为小写作为map的键
		fieldName := typ.Field(i).Name
		result[fieldName] = fieldValue
	}

	return result
}
