package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	bytes, err := os.ReadFile("E:\\RemovedD\\code\\letgoV2\\your_code\\enter.go")
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", bytes)

	fmt.Printf("%v", strings.Split(str, "\n"))

}
