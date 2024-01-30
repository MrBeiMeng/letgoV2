package func_operator

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type Arr []interface{}

func ConvertToArr(value reflect.Value) (Arr, error) {
	// Check if the value is an array
	if value.Kind() != reflect.Array && value.Kind() != reflect.Slice {
		return nil, errors.New("value is not an array or slice")
	}

	// Get the length of the array
	length := value.Len()

	// Create a new slice of interface{} to hold the elements
	arr := make(Arr, length)

	// Iterate over each element of the array
	for i := 0; i < length; i++ {
		// Get the i-th element
		element := value.Index(i)

		// Check if the type of the element matches the type of Arr
		if !element.Type().AssignableTo(reflect.TypeOf(arr).Elem()) {
			return nil, errors.New("element type does not match Arr type")
		}

		// Set the i-th element of the new Arr
		arr[i] = element.Interface()
	}

	return arr, nil
}

func (a Arr) String() string {
	strList := make([]string, 0)

	for _, anyType := range a {
		strList = append(strList, fmt.Sprintf("%v", anyType))
	}

	return "[" + strings.Join(strList, ",") + "]"
}
