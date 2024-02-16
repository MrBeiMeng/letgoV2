package util

import (
	"errors"
	"fmt"
	"io/fs"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/util/config_util"
	"path/filepath"
	"regexp"
	"strings"
)

func GetLastDirID() (error, uint32) {
	targetPath := config_util.Get("CodePlace")

	codeDirList := make([]string, 0)
	dNameList := make([]string, 0)
	filepath.WalkDir(targetPath, func(path string, d fs.DirEntry, err error) error {
		if !d.IsDir() {
			return nil
		}

		matched, _ := regexp.Match(`^ID[a-z]{4}_([a-z_-]+)$`, []byte(d.Name()))
		if !matched {
			return nil
		}

		codeDirList = append(codeDirList, path)
		dNameList = append(dNameList, d.Name())
		return nil
	})

	var maxId uint32 = 0
	var lastDir string

	for _, dirName := range dNameList {
		// 定义正则表达式
		re := regexp.MustCompile(`^ID([a-z]{4})_`)

		// 使用正则表达式查找匹配项
		match := re.FindStringSubmatch(dirName)
		if len(match) < 2 {
			continue
		}
		matchStr := match[1]
		err, id := ConvZzza2int(matchStr)
		if err != nil {
			return err, 0
		}

		if id > maxId {
			maxId = id
			lastDir = dirName
		}
	}

	logging.Info(fmt.Sprintf("查询到最后创建的文件夹[%s]ID[%d]", lastDir, maxId))
	return nil, maxId
}

// ConvZzza2int 将四位 aaaa 是 45_6975 |  1 zzzz 是0 ! 这并非十六进制
func ConvZzza2int(abc string) (error, uint32) {
	if len(abc) != 4 {
		errStr := fmt.Sprintf("字符串长度不合法")
		logging.Info(errStr)

		return errors.New(errStr), 0
	}

	var result uint32
	base := uint32('a') // 'f' 对应的数字减一作为基准 a = 97

	for _, char := range abc {
		if char < 'a' || char > 'z' {
			// 字符不在合法范围内
			errStr := fmt.Sprintf("字符不在合法范围内 [f - z].")
			logging.Info(errStr)
			return errors.New(errStr), 0
		}

		// 将字符映射为数字并累加到结果中
		result = result*26 + (25 - (uint32(char) - base))
	}

	return nil, result
}

// ConvInt2zzza 将四位 1 转为数字 fffa
func ConvInt2zzza(num uint32) (error, string) {
	// 数字范围应该是 0 到 45_6976-1
	if num < 0 || num >= 45_6976 {
		errStr := fmt.Sprintf("数字不合法")
		logging.Info(errStr)

		return errors.New(errStr), ""
	}

	var result strings.Builder
	base := uint32('a')

	for num >= 1 {
		per := num % 26
		num = num / 26
		result.WriteByte(byte(base + (25 - per)))
	}

	length := result.Len()
	for i := 0; i < 4-length; i++ {
		result.WriteByte('z')
	}

	str := result.String()

	result.Reset()
	for i := len(str) - 1; i >= 0; i-- {
		result.WriteByte(str[i])
	}

	return nil, result.String()
}
