package func_operator

import (
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"reflect"
	"strings"
)

func mustBeFunc(function interface{}) error {
	typeOfF := reflect.TypeOf(function)

	if typeOfF.Kind() != reflect.Func {
		errStr := fmt.Sprintf("传入runFunc类型[%s]并非函数，请确保传入一个函数", typeOfF.Name())
		logging.Info(errStr)
		return errors.New(errStr)
	}

	return nil
}

func sizeMustEqual(function interface{}, strParams []string) error {
	typeOfF := reflect.TypeOf(function)

	if typeOfF.NumIn() != len(strParams) {
		errStr := fmt.Sprintf("参数个数与实验函数参数个数不同，请确保测试用例正确")
		logging.Info(errStr)
	}

	return nil
}

// 使用\n来切割参数
func splitStrParams(strParams string) []string {
	strParamSlice := strings.Split(strParams, "\n")

	for i := 0; i < len(strParamSlice); i++ {

		strParamSlice[i] = strings.TrimSpace(strParamSlice[i])

	}

	return strParamSlice
}

func getFunctionParamKindSlice(function interface{}) (result []reflect.Kind) {
	typeOfFunction := reflect.TypeOf(function)

	for i := 0; i < typeOfFunction.NumIn(); i++ {
		result = append(result, typeOfFunction.In(i).Kind())
	}

	return
}

func getFunctionParamTypeSlice(function interface{}) (result []reflect.Type) {
	typeOfFunction := reflect.TypeOf(function)

	for i := 0; i < typeOfFunction.NumIn(); i++ {
		result = append(result, typeOfFunction.In(i))
	}

	return
}
