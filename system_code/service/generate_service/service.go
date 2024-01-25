package generate_service

type GenerateServiceI interface {
	GenerateFiles(param CombinedFileParams) error
}

var GenerateService GenerateServiceI

func init() {
	GenerateService = &GenerateServiceImpl{}
}
