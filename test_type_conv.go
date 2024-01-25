package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var funcCheck func(a int, b float32) float32

func typeConv(paramType []reflect.Type, paramStr string) []interface{} {

	panic("implement me")
}

func aPlusB(a int, b float32) float32 {

	return float32(a) + b
}

func main() {
	runFunc := aPlusB

	typeOf := reflect.TypeOf(runFunc)

	funcTypeOf := reflect.TypeOf(funcCheck)

	assignableTo := typeOf.AssignableTo(funcTypeOf)

	fmt.Printf("assignableTo = %+v \n", assignableTo)
	fmt.Printf("typeOf = %+v \n", typeOf)

	println("-----")

	name := funcTypeOf.Name()
	numIn := funcTypeOf.NumIn()
	numOut := funcTypeOf.NumOut()
	inputKindSlice := make([]reflect.Kind, 0)
	for i := 0; i < numIn; i++ {
		inputKindSlice = append(inputKindSlice, funcTypeOf.In(i).Kind())
	}

	outputTypeSlice := make([]reflect.Kind, 0)
	for i := 0; i < numOut; i++ {
		outputTypeSlice = append(outputTypeSlice, funcTypeOf.Out(i).Kind())
	}

	printStr := `func name[%s] in%v out%v`
	printStr = fmt.Sprintf(printStr, name, inputKindSlice, outputTypeSlice)

	println(printStr)

	println("-----")

	inputParams := make([]string, 0)
	inputParams = append(inputParams, []string{"101\n2.5", "5\n2.1"}...)

	valueOf := reflect.ValueOf(runFunc)

	for _, paramBatch := range inputParams {

		callingValueParams := make([]reflect.Value, len(inputKindSlice))

		strParams := strings.Split(paramBatch, "\n")
		// inputKindSlice

		if len(strParams) != len(inputKindSlice) {
			panic(errors.New("error"))
		}

		for i := 0; i < len(inputKindSlice); i++ {
			strParam, kindOfParam := strParams[i], inputKindSlice[i]

			tempValue := reflect.Value{}

			switch kindOfParam {

			case reflect.Float32:
				parseFloat, err := strconv.ParseFloat(strParam, 32)
				if err != nil {
					panic(err)
				}

				tempValue = reflect.ValueOf(float32(parseFloat))

			case reflect.Int:
				intS, err := strconv.Atoi(strParam)
				if err != nil {
					panic(err)
				}

				tempValue = reflect.ValueOf(intS)

			default:
				panic(errors.New(fmt.Sprintf("未处理的参数类型[%s]", kindOfParam)))
			}

			callingValueParams[i] = tempValue

		}

		// 运行方法
		funcResultValues := valueOf.Call(callingValueParams)

		for _, value := range funcResultValues {

			fmt.Printf("%+v\n", value)
		}

	}

}
