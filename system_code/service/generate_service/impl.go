package generate_service

import (
	"letgoV2/system_code/pkg/setting"
	"letgoV2/system_code/pkg/util"
	"path"
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
	basePath := setting.CodePlace
	fileContentMap := make(map[string]string)

	// 获取最后的ID，接下来生成的ID加一
	err, lastId := GetLastDirID()
	if err != nil {
		return
	}

	err, dirName := GetNewDirName(lastId+1, param.TitleSlug)
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

	err, fileContent3, fileName3 := getSnippetFileMetaDataDGo(*NewFileMetaData(dirName, param.JsonExampleTestcases))
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

	return nil
}
