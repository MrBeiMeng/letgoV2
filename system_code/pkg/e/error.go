package e

import "runtime"

type LocationError struct {
	Msg string
	Location
}

func NewLocationError(msg string, location Location) *LocationError {
	return &LocationError{Location: location, Msg: msg}
}

func NewLocationErrorHere(msg string) *LocationError {
	// 调用 runtime.Caller 获取调用栈信息
	_, file, line, _ := runtime.Caller(1)
	return &LocationError{Location: Location{
		FilePath: file, Line: uint(line),
	}, Msg: msg}
}

func (l *LocationError) GetLocation() Location {
	return l.Location
}

func (l *LocationError) Error() string {
	return l.Msg
}
