package util

import (
	"errors"
	"fmt"
	"letgoV2/system_code/pkg/logging"
	"letgoV2/system_code/pkg/logging/http_logging"
	"net/http"
	"reflect"
	"strings"
)

func HttpPost(url string, cookies string, headerMap map[string]string, requestBody string) (error, []byte) {
	client := &http.Client{}
	// ...
	req, err := http.NewRequest("POST", url, strings.NewReader(requestBody))
	if err != nil {
		panic(err)
	}

	for _, cookie := range strings.Split(cookies, ";") {
		cookie = strings.TrimSpace(cookie)
		cs := strings.Split(cookie, "=")
		if len(cs) < 2 {
			errStr := fmt.Sprintf("cookie格式不正确")
			logging.Error(errStr)
			return errors.New(errStr), nil
		}
		name := cs[0]
		value := cs[1]
		req.AddCookie(&http.Cookie{Name: name, Value: value})
	}

	for key, value := range headerMap {
		req.Header.Add(key, value)
	}

	all, err := http_logging.HttpPostLog(client, req)
	if err != nil {
		return err, nil
	}

	return nil, all
}

// CopyStructFields 把A的值复制到B上 ! 注意a是结构体，b是结构体指针
func CopyStructFields(a interface{}, b interface{}) {
	valA := reflect.ValueOf(a)
	valB := reflect.ValueOf(b).Elem()

	if valA.Kind() != reflect.Struct || valB.Kind() != reflect.Struct {
		fmt.Println("Both parameters must be structs")
		return
	}

	for i := 0; i < valA.NumField(); i++ {
		fieldA := valA.Field(i)
		fieldB := valB.FieldByName(valA.Type().Field(i).Name)

		if fieldB.IsValid() {
			// Check if the field exists in B and is not nil or zero
			if !isZero(fieldB) {
				// Check if the values are not equal
				if !reflect.DeepEqual(fieldA.Interface(), fieldB.Interface()) {
					fmt.Printf("Warning: Field %s is not equal between structs\n", valA.Type().Field(i).Name)
				}
			}
			// Copy the value from A to B
			fieldB.Set(fieldA)
		}

	}
}

// isZero checks if a field is nil or zero value
func isZero(field reflect.Value) bool {
	switch field.Kind() {
	case reflect.Ptr, reflect.Interface:
		return field.IsNil()
	default:
		zero := reflect.Zero(field.Type())
		return reflect.DeepEqual(field.Interface(), zero.Interface())
	}
}
