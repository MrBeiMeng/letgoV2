package config_util

type logStruct struct {
	Level string `yaml:"Level"`
	Dir   string `yaml:"Dir"`
}

type leetcode struct {
	Account     string
	Password    string
	Cookies     string
	ContentType string
	Origin      string
	UserAgent   string
}

type configEntity struct {
	CodePlace string    `yaml:"CodePlace"`
	Log       logStruct `yaml:"log"`
	Leetcode  leetcode  `yaml:"leetcode"`
}
