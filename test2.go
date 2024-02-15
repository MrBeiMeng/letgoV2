package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/sirupsen/logrus"
)

func main() {
	// 设置logrus输出到标准输出
	logrus.SetOutput(os.Stdout)

	// 设置日志级别为Error
	logrus.SetLevel(logrus.ErrorLevel)

	// 模拟发生错误的函数
	err := someFunction()
	if err != nil {
		// 获取错误发生的文件名和行号
		_, file, line, ok := runtime.Caller(0)
		if ok {
			// 使用 //line 注释指定文件和行号
			fmt.Printf("%s:%d\n", file, line)

			// 添加地址信息到日志
			logrus.WithFields(logrus.Fields{
				"file": file,
				"line": line,
			}).Error("An error occurred")
		} else {
			// 如果无法获取地址信息，则只记录错误信息
			logrus.Error("An error occurred")
		}
	}
}

func someFunction() error {
	// 模拟发生错误
	return fmt.Errorf("something went wrong")
}
