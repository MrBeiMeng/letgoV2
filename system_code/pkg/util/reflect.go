package util

import "reflect"

func StructToMap(obj interface{}) map[string]interface{} {
	// 获取结构体的反射值
	val := reflect.ValueOf(obj)

	// 如果传入的不是结构体对象，则返回空map
	if val.Kind() != reflect.Struct {
		return nil
	}

	// 创建一个map存储转换后的键值对
	result := make(map[string]interface{})

	// 获取结构体的类型
	typ := reflect.TypeOf(obj)

	// 遍历结构体的字段并添加到map中
	for i := 0; i < val.NumField(); i++ {
		// 获取字段的值
		fieldValue := val.Field(i).Interface()

		// 获取字段的名称，并转换为小写作为map的键
		fieldName := typ.Field(i).Name
		result[fieldName] = fieldValue
	}

	return result
}
