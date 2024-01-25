package util

import (
	"fmt"
	"io/ioutil"
	"letgoV2/system_code/pkg/logging"
	"os"
)

func CreateAndWriteToFile(fileName string, content string) error {
	//  检查文件名和内容是否为空，如果是则返回错误
	if fileName == "" || content == "" {
		return fmt.Errorf("文件名和内容不能为空")
	}

	//  使用ioutil.WriteFile创建文件并写入内容，如果出现错误则返回错误
	bitContent := []byte(content)
	err := ioutil.WriteFile(fileName, bitContent, 0644)
	if err != nil {
		return fmt.Errorf("无法写入文件：%v", err)
	}

	//  返回nil表示函数执行成功，没有错误发生
	logging.Info(fmt.Sprintf("创建了文件[%s]并写入了[%d]bite内容", fileName, len(bitContent)))
	return nil
}

func DeleteFileOrFolder(path string) error {
	//  检查路径是否为空，如果是则返回错误
	if path == "" {
		return fmt.Errorf("路径不能为空")
	}

	//  获取文件/文件夹信息，以检查存在性
	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("文件或文件夹不存在：%s", path)
	}

	//  根据文件/文件夹类型执行删除操作
	if fileInfo.IsDir() {
		// 如果是文件夹，使用os.RemoveAll删除
		err := os.RemoveAll(path)
		if err != nil {
			return fmt.Errorf("无法删除文件夹：%v", err)
		}
		logging.Info(fmt.Sprintf("删除了文件夹[%s]", path))
	} else {
		// 如果是文件，使用os.Remove删除
		err := os.Remove(path)
		if err != nil {
			return fmt.Errorf("无法删除文件：%v", err)
		}
		logging.Info(fmt.Sprintf("删除了文件[%s]", path))
	}

	//  返回nil表示函数执行成功，没有错误发生

	return nil
}

func CreateFolder(folderPath string) error {
	// 检查文件夹路径是否为空，如果是则返回错误
	if folderPath == "" {
		return fmt.Errorf("文件夹路径不能为空")
	}

	// 使用os.Stat检查文件夹是否已经存在
	if _, err := os.Stat(folderPath); err == nil {
		// 文件夹已经存在，记录警告信息
		logging.Warn(fmt.Sprintf("文件夹[%s]已经存在", folderPath))
		return nil
	} else if !os.IsNotExist(err) {
		// 发生了其他错误，返回错误信息
		return fmt.Errorf("无法检查文件夹：%v", err)
	}

	// 使用os.Mkdir创建文件夹，第二个参数是文件夹权限，这里使用默认权限 0755
	err := os.Mkdir(folderPath, 0755)
	if err != nil {
		return fmt.Errorf("无法创建文件夹：%v", err)
	}

	// 返回nil表示函数执行成功，没有错误发生
	logging.Info(fmt.Sprintf("创建了文件夹[%s]", folderPath))
	return nil
}
