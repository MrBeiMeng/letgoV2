package cmd_params

type DownParam struct {
	ID        string // 题目id
	TitleSlug string // 题目Slug
}

type RemoveParam struct {
	DirId   string
	Confirm bool
}

type SetParam struct {
	Set     string
	Show    string
	Cookies string
	Confirm bool
}
