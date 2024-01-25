package main

import (
	"fmt"
	"reflect"
)

type Test struct {
	Name string `replace:"name" json:"name"`
	Age  int
}

func main() {

	testP := Test{
		"tom",
		12,
	}

	//tV := reflect.ValueOf(testP)
	tT := reflect.TypeOf(testP)

	for i := 0; i < tT.NumField(); i++ {
		//fieldValue := tV.Field(i)
		fieldType := tT.Field(i)
		fmt.Printf("tag = %s \n", fieldType.Tag)
	}

}
