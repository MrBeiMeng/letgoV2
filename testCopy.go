package main

import (
	"fmt"
	"letgoV2/system_code/pkg/util"
)

type A struct {
	Name string
	Age  int
}

type B struct {
	Name string
	Age  int
	Addr string
}

func main() {

	a := A{
		Name: "Tom",
		Age:  18,
	}

	b := B{
		Addr: "上海",
	}

	util.CopyStructFields(a, &b)

	fmt.Printf("%#v", b)

}
