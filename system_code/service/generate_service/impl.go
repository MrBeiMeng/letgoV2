package generate_service

import (
	"fmt"
	"io/ioutil"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/util"
	"letgoV2/system_code/pkg/util/config_util"
	"os"
	"path"
	"strings"
)

type GenerateServiceImpl struct {
}

// GenerateFiles 注释步骤
// 有个计数器
// 生成文件夹名
// 生成 code.go 文件
// 生成 code_test.go 文件
// 生成 meta_data.go 文件
func (g *GenerateServiceImpl) GenerateFiles(param CombinedFileParams) (err error) {
	basePath := config_util.Get("CodePlace")
	fileContentMap := make(map[string]string)

	// 获取最后的ID，接下来生成的ID加一
	err, lastId := util.GetLastDirID()
	if err != nil {
		return
	}

	newDirId := lastId + 1
	err, dirName := GetNewDirName(newDirId, param.TitleSlug)
	if err != nil {
		return
	}

	baseDirPath := path.Join(basePath, dirName)

	err, fileContent1, fileName1 := getSnippetFileCodeDGo(*newFileCodeParam(dirName, param.GoCodeSnippet))
	if err != nil {
		return
	}
	fileContentMap[fileName1] = fileContent1

	err, fileContent2, fileName2 := getSnippetFileCodeTestDGo(*newFileCodeTestParam(dirName, param.GoCodeSnippet))
	if err != nil {
		return
	}
	fileContentMap[fileName2] = fileContent2

	err, fileContent3, fileName3 := getSnippetFileMetaDataDGo(*NewFileMetaData(dirName, param.JsonExampleTestcases, param.GoCodeSnippet, newDirId))
	if err != nil {
		return
	}
	fileContentMap[fileName3] = fileContent3

	err, fileContent4, fileName4 := getSnippetFileReadMeEnDMd(*newFileReadMeEnParam(param))
	if err != nil {
		return
	}
	fileContentMap[fileName4] = fileContent4

	err, fileContent5, fileName5 := getSnippetFileReadMeZhDMd(*newFileReadMeZhParam(param))
	if err != nil {
		return
	}
	fileContentMap[fileName5] = fileContent5

	err, funName := getOneFuncName(param.GoCodeSnippet)
	if err != nil {
		logging.Error(err)
	}

	// 创建tests文件夹
	dir, _ := os.Getwd()
	testDir := path.Join(dir, fmt.Sprintf("system_code/pkg/tests/%s", CamelToSnake(funName)))
	err = util.CreateFolder(testDir)
	if err != nil {
		logging.Error(err)
	}

	//if _, err := os.Stat(path.Join(testDir, "test.go")); err != nil {
	//	createTestDGo(funName, err, testDir)
	//}

	// 创建文件保存内容
	err = util.CreateFolder(baseDirPath)
	if err != nil {
		return
	}

	for fileName, fileContent := range fileContentMap {
		filePath := path.Join(baseDirPath, fileName)
		err = util.CreateAndWriteToFile(filePath, fileContent)
		if err != nil {
			return
		}
	}

	filePath := path.Join(basePath, "enter.go")
	fileContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("无法读取文件:", err)
		return
	}

	fileContent = []byte(strings.ReplaceAll(string(fileContent), "\n)", fmt.Sprintf("\n\t_ \"letgoV2/your_code/%s\"\n)", dirName)))

	err = util.CreateAndWriteToFile(filePath, string(fileContent))
	if err != nil {
		panic(err)
	}

	return nil
}

func createTestDGo(funName string, err error, testDir string) {
	fileContent6 := `package ${CamelFuncName}

import "letgoV2/system_code/service/code_handle_service/code_handle_params"

var (
	Tests = []code_handle_params.Test{ // 如果您愿意，感谢您将自己编写的Test提交到远程仓库中，这会帮到很多人。
		//{TestStr: "", ResultChecker: nil,ShowWhenErr: "you made a mistake --by githubName"},
	}
)
`

	fileContent6 = strings.ReplaceAll(fileContent6, "${CamelFuncName}", CamelToSnake(funName))

	err = util.CreateAndWriteToFile(path.Join(testDir, "test.go"), fileContent6)
	if err != nil {
		logging.Error(err)
	}
}
