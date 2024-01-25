package main

import (
	"fmt"
	"letgoV2/system_code/service/generate_service"
)

func main() {

	str := "zzzz"

	fmt.Printf("str = %s\n", str)
	err, num := generate_service.ConvZzza2int(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("num = %d\n", num)

	//num := 45_6976 - 1

	fmt.Printf("num = %d\n", num)
	err, strS := generate_service.ConvInt2zzza(uint32(num))
	if err != nil {
		panic(err)
	}
	fmt.Printf("strS = %s\n", strS)

}
