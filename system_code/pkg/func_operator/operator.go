package func_operator

import (
	"reflect"
	"strings"
)

// RunFunc 运行代码
func RunFunc(function interface{}, strParams string) (error, string) {

	return runFunc(function, splitStrParams(strParams))
}

// runFunc 运行代码
// function test
// 参数个数检测
func runFunc(function interface{}, strParams []string) (error, string) {
	err := mustBeFunc(function)
	if err != nil {
		return err, ""
	}

	err = sizeMustEqual(function, strParams)
	if err != nil {
		return err, ""
	}

	callValueSlice := make([]reflect.Value, 0)
	typeSlice := getFunctionParamTypeSlice(function)
	for i := range typeSlice {
		paramType := typeSlice[i]
		strParam := strParams[i]

		err, value := convStr2TypeValue(strParam, paramType)
		if err != nil {
			return err, ""
		}

		callValueSlice = append(callValueSlice, value)
	}

	valueOfFunction := reflect.ValueOf(function)
	resultValueSlice := valueOfFunction.Call(callValueSlice)

	outputStr := strings.Builder{}
	for _, resultValue := range resultValueSlice {
		err, resStr := convTypeValue2Str(resultValue)
		if err != nil {
			return err, ""
		}

		outputStr.WriteString(resStr)
		//outputStr.WriteString("\n")
	}

	return nil, outputStr.String()
}
