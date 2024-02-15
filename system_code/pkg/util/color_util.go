package util

import "fmt"

const (
	RED    = 31
	GREEN  = 32
	YELLOW = 33
	BLUE   = 36
	GRAY   = 37
)

// SetColor 设置颜色 为了避免颜色保存到 通过两部分实现
// 1 调用SetColor方法设置颜色，
// 2 调用Paint方法进行上色
// 3
func SetColor(content string, color int) string {

	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color, content)
}

//func SetColorCase(content string, caseFunc func(value interface{}) string) string {
//
//}
