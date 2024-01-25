package func_operator

import (
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"reflect"
	"strconv"
)

// 将字符串类型转为golang类型
func convStr2TypeValue(strParam string, valueOf reflect.Type) (error, reflect.Value) {
	// 难点在于嵌套的类型，比如二维数组
	result := reflect.Value{}

	switch valueOf.Kind() {
	case reflect.String:
		result = reflect.ValueOf(strParam)
	case reflect.Int:
		finalInt, err := strconv.Atoi(strParam)
		if err != nil {
			return err, reflect.Value{}
		}
		result = reflect.ValueOf(finalInt)
	case reflect.Float64:
		finalFloat, err := strconv.ParseFloat(strParam, 64)
		if err != nil {
			return err, reflect.Value{}
		}
		result = reflect.ValueOf(finalFloat)
	case reflect.Bool:
		finalBool, err := strconv.ParseBool(strParam)
		if err != nil {
			return err, reflect.Value{}
		}
		result = reflect.ValueOf(finalBool)
	case reflect.Slice, reflect.Array:

		panic("implement me")

		//// 处理切片和数组的情况，递归调用 convStr2TypeValue
		//elemType := valueOf.Elem()
		//elemValue := reflect.New(elemType).Elem()
		//// 这里需要根据具体情况调整，例如处理逗号分隔的字符串转为切片
		//// 这里仅作示例，具体实现需要根据需求修改
		//// elemValue = convertStringToSlice(strParam, elemType)
		//result = reflect.MakeSlice(valueOf, 1, 1)
		//result.Index(0).Set(elemValue)

	default:
		errStr := fmt.Sprintf("未处理的类型转换")
		logging.Error(errStr)
		return errors.New(errStr), reflect.Value{}
	}

	return nil, result
}

func convTypeValue2Str(value reflect.Value) (error, string) {
	// 尝试使用 String() 方法获取字符串表示
	if value.Kind() == reflect.String {
		return nil, value.String()
	}

	// 如果底层类型不是字符串，进行适当的类型转换
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return nil, strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return nil, strconv.FormatUint(value.Uint(), 10)
	case reflect.Float32, reflect.Float64:
		return nil, strconv.FormatFloat(value.Float(), 'f', -1, 64)
	case reflect.Bool:
		return nil, strconv.FormatBool(value.Bool())
	default:
		// 如果无法处理的类型，返回空字符串或者进行其他处理
		panic("implement me")
	}
}
