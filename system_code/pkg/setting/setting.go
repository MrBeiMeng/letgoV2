package setting

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"path"
)

var (
	Cfg          Config
	LeetCodeConf LeetCodeConfT
	LogDir       string
	CodePlace    string
	NotRootStart bool = false
)

type Config struct {
	LeetCodeConf LeetCodeConfT `yaml:"leetcode-conf"`
	LogAbout     struct {
		LogDir string `yaml:"log-dir"`
	} `yaml:"log-about"`
	CodePlace string `yaml:"code-place"`
}

type LeetCodeConfT struct {
	Cookies string `yaml:"cookies"`

	HeaderMap map[string]string

	RandomUuid  string `yaml:"random-uuid"`
	ContentType string `yaml:"content-type"`
	Origin      string `yaml:"origin"`
	XCsrftoken  string `yaml:"x-csrftoken"`
	UserAgent   string `yaml:"user-agent"`
}

func init() {
	loadCfgFromYaml()

	exploreLeetCodeConf()
	exploreLogDir()
	exploreCodePlace()
}

func exploreCodePlace() {
	pwd, _ := os.Getwd()
	CodePlace = path.Join(pwd, Cfg.CodePlace)
	if CodePlace == "" {
		fmt.Println("CodePlace is \"\"")
	}
}

func loadCfgFromYaml() {
	// 打开 YAML 文件
	file, err := os.Open("system_code/conf/conf.yaml")
	if err != nil {
		// 文件无法打开，也可能是your_code下面的方法调用导致的
		file, err = os.Open("../../system_code/conf/conf.yaml")
		if err != nil {
			panic(fmt.Sprintf("Error opening file: %v", err))
		}
		NotRootStart = true
	}
	defer file.Close()

	// 创建解析器
	decoder := yaml.NewDecoder(file)

	// 解析 YAML 数据
	err = decoder.Decode(&Cfg)
	if err != nil {
		panic(fmt.Sprintf("Error decoding YAML:%v", err))
	}
}

func exploreLogDir() {
	LogDir = Cfg.LogAbout.LogDir
	if LogDir == "" {
		fmt.Println("LogDir is \"\"")
	}
}

func exploreLeetCodeConf() {
	headerMap := make(map[string]string)

	headerMap["content-type"] = Cfg.LeetCodeConf.ContentType
	headerMap["origin"] = Cfg.LeetCodeConf.Origin
	headerMap["user-agent"] = Cfg.LeetCodeConf.UserAgent

	Cfg.LeetCodeConf.HeaderMap = headerMap

	LeetCodeConf = Cfg.LeetCodeConf
}
