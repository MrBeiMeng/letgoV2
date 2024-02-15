package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

func init() {
	// 设置logrus输出到标准输出
	logrus.SetOutput(os.Stdout)

	// 如果连接了TTY，启用颜色
	if IsTerminal() {
		logrus.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
	} else {
		logrus.SetFormatter(&logrus.JSONFormatter{})
	}
}

// IsTerminal 判断是否连接了TTY
func IsTerminal() bool {
	fileInfo, err := os.Stdout.Stat()
	if err != nil {
		return false
	}
	return (fileInfo.Mode() & os.ModeCharDevice) != 0
}

func main() {
	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("A group of walrus emerges from the ocean")

	logrus.WithFields(logrus.Fields{
		"animal": "walrus",
		"size":   10,
	}).Error("A group of walrus emerges from the ocean")

	logrus.Warn("This is warning!")
}
