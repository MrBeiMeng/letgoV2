package e

type Location struct {
	FilePath string
	Line     uint
}

// CustomError 是自定义的错误类型，包含文件地址和行数
type CustomError interface {
	// GetLocation 获取错误发生的地址
	GetLocation() Location
}
