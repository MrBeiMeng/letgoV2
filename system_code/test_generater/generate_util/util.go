package generate_util

import (
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
)

func convStr2Int(str string) int {

	if strings.Contains(str, "**") {
		nums := strings.Split(str, "**")

		a, b := 0, 0
		if len(nums) > 1 {
			a, _ = strconv.Atoi(nums[0])
			b, _ = strconv.Atoi(nums[1])
		}

		var result int = 1

		for i := 0; i < b; i++ {
			result *= int(a)
		}

		return result
	}

	num, _ := strconv.Atoi(str)
	return int(num)
}

func GetNum(left, right string, seed *int64) int {
	rand.Seed(int64(*seed))
	*seed += 1

	return rand.Intn(convStr2Int(right)-convStr2Int(left)+1) + convStr2Int(left)
}

func Arr2Str(arr interface{}) string {
	strBuilder := strings.Builder{}
	strBuilder.WriteString("[")

	valueOf := reflect.ValueOf(arr)

	for i := 0; i < valueOf.Len(); i++ {
		strBuilder.WriteString(fmt.Sprintf("%v", valueOf.Index(i)))

		if i < valueOf.Len()-1 {
			strBuilder.WriteString(",")
		}
	}

	strBuilder.WriteString("]")

	return strBuilder.String()
}
