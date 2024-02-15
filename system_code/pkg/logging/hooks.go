package logging

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

// ColorfulConsoleHook 自定义的Console hook，用于控制台输出带颜色的日志
type ColorfulConsoleHook struct {
	formatter logrus.Formatter
}

// NewColorfulConsoleHook 创建一个新的ColorfulConsoleHook实例
func NewColorfulConsoleHook(formatter logrus.Formatter) *ColorfulConsoleHook {
	return &ColorfulConsoleHook{
		formatter: formatter,
	}
}

// Fire 实现Hook接口中的Fire方法，用于处理日志事件
func (hook *ColorfulConsoleHook) Fire(entry *logrus.Entry) error {

	bytes, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = os.Stdout.Write(bytes)
	return err
}

// Levels 实现Hook接口中的Levels方法，指定hook处理的日志级别
func (hook *ColorfulConsoleHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// FileHook 自定义的File hook，用于输出到文件的日志
type FileHook struct {
	writer    io.Writer
	formatter logrus.Formatter
}

// NewFileHook 创建一个新的FileHook实例
func NewFileHook(writer io.Writer, formatter logrus.Formatter) *FileHook {
	return &FileHook{
		writer:    writer,
		formatter: formatter,
	}
}

// Fire 实现Hook接口中的Fire方法，用于处理日志事件
func (hook *FileHook) Fire(entry *logrus.Entry) error {
	bytes, err := hook.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = hook.writer.Write(bytes)
	return err
}

// Levels 实现Hook接口中的Levels方法，指定hook处理的日志级别
func (hook *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
