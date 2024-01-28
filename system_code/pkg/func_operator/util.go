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

// PeelOffALayer
// 去掉一层外壳 ，例如 func("[[1,2,3],[2,3,4]]") = "[1,2,3]","[2,3,4]"
// 前提是数组或切片类型的字符串参数
func PeelOffALayer(strParams string) (error, []string) {
	if strParams[0] != '[' {
		errStr := fmt.Sprintf("字符串开头非\"[\"")
		logging.Info(errStr)
		return errors.New(errStr), nil
	}
	strParams = strParams[1:]

	if strParams[len(strParams)-1] != ']' {
		errStr := fmt.Sprintf("字符串结尾非\"]\"")
		logging.Info(errStr)
		return errors.New(errStr), nil
	}
	strParams = strParams[:len(strParams)-1]

	result := make([]string, 0)

	isALayer := !strings.Contains(strParams, "[")
	if isALayer {
		return nil, []string{strParams}
	}

	countDeep := 0
	builder := strings.Builder{}
	for _, subRune := range strParams {
		if countDeep > 0 {
			builder.WriteRune(subRune)
		}
		if subRune == '[' {
			if countDeep > 0 {
				builder.WriteRune(subRune)
			}
			countDeep += 1
			continue
		}
		if subRune == ']' {
			countDeep -= 1
			if countDeep == 0 {
				result = append(result, builder.String())
				builder.Reset()
				continue
			}
		}
	}

	if countDeep != 0 {
		errStr := fmt.Sprintf("非法参数:%s", strParams)
		logging.Info(errStr)
		return errors.New(errStr), nil
	}

	return nil, result
}
