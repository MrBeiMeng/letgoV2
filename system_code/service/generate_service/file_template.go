package generate_service

import (
	"letgoV2/system_code/pkg/logging"
)

func getSnippetFileCodeDGo(param fileCodeParam) (err error, snippet, fileName string) {
	fileName = "code.go"
	codeTemplate := `
package ${packageName}
import (
	_ "fmt"
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

	snippet = addImportIfNeed(snippet)
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

			logging.Info(fmt.Sprintf("longestPalindrome(%s) = %v ", sampleTest, result))
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

import (
	"letgoV2/system_code/service/code_handle_service"
	"letgoV2/system_code/service/code_handle_service/code_handle_params"
)


var (
	// sampleTests 是为了您在编写函数时debug
	sampleTests = []string{
		${sampleTests}
	}
)

func init() {

	// 与sampleTests不同，这里的Test将在您使用命令行RUN时被调用，
	// tests 是为了写好函数后统一测试
	tests := []code_handle_params.Test{
		//{TestStr: "", CorrectResult: nil,ShowWhenErr: "you made a mistake --by githubName"},
		${tests}
	}

	// 当TestStr在 ${underlineFuncName}.Tests 已经存在时 tests中写的测试用例将会被 ${underlineFuncName}.Tests 中的用例覆盖
	_ = code_handle_service.CodeHandleService.SignIn("zzzz", ${funcName}, tests)
	_ = code_handle_service.CodeHandleService.SignInTestFile("zzzz", "system_code/pkg/tests/${underlineFuncName}")

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
