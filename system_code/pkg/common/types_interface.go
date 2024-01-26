package common

type LoadI interface {
	LoadFromStr(str string) error
}
