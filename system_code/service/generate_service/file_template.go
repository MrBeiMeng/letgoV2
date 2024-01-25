package generate_service

import "letgoV2/system_code/pkg/logging"

func getSnippetFileCodeDGo(param fileCodeParam) (err error, snippet, fileName string) {
	fileName = "code.go"
	codeTemplate := `
package ${packageName}

import (
	_ "letgoV2/system_code/pkg/common"
)


${goCodeSnippet}

`

	err, snippet = replaceStructParam(codeTemplate, param)

	// 完善代码
	err, snippet = UpdateGoCodeSnippetByAst(snippet)
	if err != nil {
		logging.Error(err)
		return
	}

	return
}

func getSnippetFileCodeTestDGo(param fileCodeTestParam) (err error, snippet, fileName string) {
	fileName = "code_test.go"
	codeTestTemplate := `
package ${packageName}

import (
	"fmt"
	"letgoV2/system_code/pkg/func_operator"
	"letgoV2/system_code/pkg/logging"
	"testing"
)

func Test_${funcName}(t *testing.T) {

	for i, sampleTest := range sampleTests { // 你可以从meta_data.go 中找到测试集合
		t.Run(fmt.Sprintf("CASE %d", i+1), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("发生了panic:", r)
					logging.Error(r)
					t.Error(r)
				}
			}()

			err, result := func_operator.RunFunc(${funcName}, sampleTest)
			if err != nil {
				logging.Error(err)
				t.Error(err.Error())
			}

			fmt.Printf("${funcName}(%s) = %v \n", sampleTest, result)
		})
	}

}

`

	err, snippet = replaceStructParam(codeTestTemplate, param)
	return
}

func getSnippetFileMetaDataDGo(param FileMetaData) (err error, snippet, fileName string) {
	fileName = "meta_data.go"
	codeMetaDataTemplate := `
package ${packageName}

var (
	sampleTests []string
)

func init() {
	sampleTests = make([]string, 0)

	sampleTests = append(sampleTests, []string{
		${sampleTests},
	}...)

}

`

	err, snippet = replaceStructParam(codeMetaDataTemplate, param)
	return
}

func getSnippetFileReadMeEnDMd(param fileReadMeEnParam) (err error, snippet, fileName string) {
	fileName = "README-en.md"
	codeReadMeEnDMd := `
<hr style="background:#ffd04c;margin: 0 200px;height:18px;border-radius:5px">

# ${englishTitle}

[<span style="font-weight:bold;font-size:14px">中文</span>](./README-zh.md) | [<span style="font-weight:bold;font-size:14px">ENGLISH 👈</span>](./README-en.md)

<span style="font-weight:bold;font-size:14px">Difficulty</span> ${difficultySpan} <span style="font-weight:bold;font-size:14px">URL</span> [${questionId}.${englishTitle}](${url})

${englishContent}

<hr style="background:#ffd04c;margin: 0 60px">


<span style="font-weight:bold;font-size:14px">passing rate: ${acRate}</span>  <span style="font-weight:bold;font-size:14px" alt="${tagList}">Click to view tags</span> 

`

	err, snippet = replaceStructParam(codeReadMeEnDMd, param)
	return
}

func getSnippetFileReadMeZhDMd(param fileReadMeZhParam) (err error, snippet, fileName string) {
	fileName = "README-zh.md"
	codeReadMeZhDMd := `
<hr style="background:#ffd04c;margin: 0 200px;height:18px;border-radius:5px">


# ${translatedTitle}

[<span style="font-weight:bold;font-size:14px">中文 👈</span>](./README-zh.md) | [<span style="font-weight:bold;font-size:14px">ENGLISH</span>](./README-en.md)

<span style="font-weight:bold;font-size:14px">难度</span> ${difficultySpan}  <span style="font-weight:bold;font-size:14px">地址</span> [${questionId}.${translatedTitle}](${url})

${translatedContent}

<hr style="background:#ffd04c;margin: 0 60px">


<span style="font-weight:bold;font-size:14px">通过率：${acRate}</span>  <span style="font-weight:bold;font-size:14px" alt="${tagList}">点击查看标签</span>

`

	err, snippet = replaceStructParam(codeReadMeZhDMd, param)
	return
}
