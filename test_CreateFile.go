package main

import (
	"fmt"
	"io/ioutil"
	"letgoV2/system_code/pkg/util"
	"strings"
)

func main() {

	filePath := "E:\\RemovedD\\code\\letgoV2\\your_code\\enter.go"
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return
	}

	fileContent = []byte(strings.ReplaceAll(string(fileContent), "\n)", "\n\t_ \"letgoV2/your_code/IDzzyy_two_sum\"\n)"))

	err = util.CreateAndWriteToFile(filePath, string(fileContent))
	if err != nil {
		panic(err)
	}
}
