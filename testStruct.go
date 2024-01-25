package main

import (
	"fmt"
	"reflect"
)

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func main() {

	LinkList := LinkNode{Val: 3, Next: &LinkNode{2, &LinkNode{1, nil}}}

	typeOf := reflect.TypeOf(LinkList)

	fmt.Printf("Kind = %s | Name = %s \n", typeOf.Kind(), typeOf.Name())

}
