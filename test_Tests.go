package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func main() {
	testStr := "[\\\"[[1,4,5],[1,3,4],[2,6]]\\\", \\\"[]\\\", \\\"[[]]\\\"]"

	tests := make([]string, 0)

	trimed := strings.ReplaceAll(testStr, "\\\"", "\"")
	err := json.Unmarshal([]byte(trimed), &tests)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\"%+v\"", tests)

}
