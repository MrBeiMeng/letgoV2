package main

import (
	"fmt"
	"letgoV2/system_code/pkg/func_operator"
)

func main() {

	strParam := "[[[1,2,3,4]],[[3,4,5]]]"

	err, strings := func_operator.PeelOffALayer(strParam)
	if err != nil {
		panic(err)
	}

	for i, s := range strings {
		fmt.Printf("%d=\"%s\"\n", i, s)
	}

}
