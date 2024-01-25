package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	var tempMap map[string]string

	tempMap = make(map[string]string)

	tempMap["yes"] = "no"

	marshal, err := json.Marshal(tempMap)
	if err != nil {
		return

	}

	println(marshal)
	fmt.Printf("%v", string(marshal))

}
