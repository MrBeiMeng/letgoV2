package IDzzzz_two_sum

import (
	"letgoV2/system_code/pkg/tests/two_sum"
	"letgoV2/system_code/service/code_handle_service"
	"letgoV2/system_code/service/code_handle_service/code_handle_params"
)

var (
	// sampleTests 是为了您在编写函数时debug
	sampleTests = []string{
		"[2,7,11,15]\n9", "[3,2,4]\n6", "[3,3]\n6",
	}
)

func init() {

	// 与sampleTests不同，这里的Test将在您使用命令行RUN时被调用，
	// tests 是为了写好函数后统一测试
	tests := []code_handle_params.Test{
		//{TestStr: "", CorrectResult: nil,ShowWhenErr: "you made a mistake --by githubName"},
		{TestStr: "[2,7,11,15]\n9", CorrectResult: nil},
		{TestStr: "[3,2,4]\n6", CorrectResult: nil},
		{TestStr: "[3,3]\n6", CorrectResult: nil},
	}

	// 当TestStr在 two_sum.Tests 已经存在时 tests中写的测试用例将会被 two_sum.Tests 中的用例覆盖
	_ = code_handle_service.CodeHandleService.SignIn("zzzz", twoSum, append(tests, two_sum.Tests...))
}
