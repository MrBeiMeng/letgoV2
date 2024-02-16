package config_util

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type DataMap map[string]interface{}

func (d *DataMap) Fields(field string) *DataMap {

	for fieldStr, value := range *d {
		if strings.EqualFold(fieldStr, field) {
			data := value.(DataMap)

			return (*DataMap)(&data)
		}

	}

	//m2 := make(map[string]interface{})
	//m := &m2
	m3 := make(map[string]interface{})
	(*d)[field] = &m3

	return (*DataMap)(((*d)[field]).(*map[string]interface{}))
}

func (d *DataMap) Set(key, value string) {

	(*d)[key] = value
}

func (d *DataMap) Get(key string) string {
	result := (*d)[key]
	if result == nil {
		panic(fmt.Sprintf("没有此key:%s", key))
	}

	return result.(string)
}

type configUtil struct {
	savePath string
	data     *DataMap
}

func newConfigUtil(configPath string) *configUtil {
	self := &configUtil{}
	self.savePath = configPath
	self.Load()

	return self
}

func (c *configUtil) GetData() *DataMap {
	c.initIfEmpty()

	return c.data
}

func (c *configUtil) initIfEmpty() {
	if c.data == nil {
		c.data = (*DataMap)(new(map[string]interface{}))
	}
}

func (c *configUtil) Save() {

	marshal, err := yaml.Marshal(*(c.data))
	if err != nil {
		panic(err.Error())
	}

	// 写入到文件
	err = ioutil.WriteFile(c.savePath, marshal, 0644)
	if err != nil {
		log.Fatalf("写入文件出错: %v", err)
	}

	fmt.Println("YAML 文件已更新并保存")
}

func (c *configUtil) Load() *DataMap {
	c.initIfEmpty()

	file, err := os.Open(c.savePath)
	if err != nil {
		file2, err2 := os.Open(filepath.Join("../../", c.savePath))
		if err2 != nil {
			panic(err2.Error())
		}
		file = file2
	}

	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err.Error())
	}
	err = yaml.Unmarshal(all, c.data)

	return c.data
}

func Save() {
	Config.Save()
}

func Fields(field string) *DataMap {

	return Config.data.Fields(field)
}

func Set(key, value string) {
	Config.data.Set(key, value)
	Config.Save()
}

func Get(key string) string {
	return Config.data.Get(key)
}

var (
	Config = newConfigUtil("system_code/conf/global_config.yaml")

	// 自动注册解析 yaml 文件
)

func init() {

}
