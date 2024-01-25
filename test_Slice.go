package main

import "fmt"

func returnSlice() (result []string) {
	result = append(result, "yes")
	return
}

func main() {

	fmt.Printf("%+v", returnSlice())

}
