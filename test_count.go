package main

import (
	"fmt"
	"io/fs"
	"letgoV2/system_code/service/generate_service"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func test() {
	var files []string

	dirSet := make(map[string]struct{})

	targetPath := "E:\\RemovedD\\code\\letgoV2\\your_code"

	err := filepath.Walk(targetPath, func(pathStr string, info os.FileInfo, err error) error {

		dir := filepath.Dir(pathStr)

		s := strings.Split(dir, "\\")
		pwDir := s[len(s)-1]

		//println(pwDir)

		matched, err := regexp.Match(`^ID[a-z]{4}_([a-z_-]+)$`, []byte(pwDir))
		if err != nil {
			println(err.Error())
		}
		if !matched {
			return nil
		}

		if _, ok := dirSet[dir]; !ok {
			dirSet[dir] = struct{}{}
		}

		files = append(files, pathStr)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for dir, _ := range dirSet {
		println(dir)
	}

	println("------")

	for _, file := range files {
		fmt.Println(file)
	}
}

func oldMain() {
	targetPath := "E:\\RemovedD\\code\\letgoV2\\your_code"

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

	for _, str := range dNameList {
		// 定义正则表达式
		re := regexp.MustCompile(`^ID([a-z]{4})_`)

		// 使用正则表达式查找匹配项
		match := re.FindStringSubmatch(str)
		matchStr := match[1]
		println(matchStr)

		count := 0

		for i := range matchStr {
			char := matchStr[i]
			if char == 'f' {
				count += 0
				continue
			}
			count += int(char - 96)
		}

		println(count)
	}
}

func main() {
	err, num := generate_service.GetLastDirID()
	if err != nil {
		panic(err)
	}

	fmt.Printf("last Id = %d\n", num)

	err, s := generate_service.GetNewDirName(num+1, "test-sshg-sg-s")
	if err != nil {
		panic(err)
	}

	fmt.Printf("dirName = %s\n", s)
}
