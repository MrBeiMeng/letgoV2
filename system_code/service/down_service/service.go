package down_service

type DownServiceI interface {
	DownById(id string) error
	DownByTitleSlug(titleSlug string) error
}

var DownService DownServiceI

func init() {
	DownService = &DownServiceImpl{}
}
