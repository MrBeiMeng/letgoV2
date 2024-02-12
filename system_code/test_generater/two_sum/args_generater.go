package main

import (
	"fmt"
	"letgoV2/system_code/test_generater/generate_util"
	"os"
	"time"
)

// 2 <= nums.length <= 10**4
// -10**9 <= nums <= 10**9
// -10**9 <= target <= 10**9
// nums\ntarget
func main() {
	seed := time.Now().Unix()
	testSize := 10

	file, err := os.OpenFile("system_code/pkg/tests/two_sum/tests_sample.json", os.O_CREATE, 511)
	if err != nil {
		panic(err.Error())
	}

	file.WriteString("[\n")

	for i := 0; i < testSize; i++ {
		templateStr := getTemplate(&seed)

		file.WriteString(templateStr)
		if i < testSize-1 {
			file.WriteString(",\n")
		}

		println(i)
		//time.Sleep(time.Second)
	}

	file.WriteString("\n]\n")

	file.Close()

}

func getTemplate(seed *int64) string {

	//length := generate_util.GetNum("2", "10**4")
	length := 9
	target := generate_util.GetNum("-10**9", "10**9", seed)
	a := generate_util.GetNum("-10**9", "10**9", seed)
	b := target - a
	//println(length)

	nums := make([]int, 0)
	for i := 0; i < length; i++ {
		numB := generate_util.GetNum("-10**9", "10**9", seed)
		for target-numB == a {
			numB = generate_util.GetNum("-10**9", "10**9", seed)
		}

		nums = append(nums, numB)
	}

	indexA := generate_util.GetNum("0", fmt.Sprintf("%d", length-1), seed)
	indexB := generate_util.GetNum("0", fmt.Sprintf("%d", length-1), seed)
	for indexB == indexA {
		indexB = generate_util.GetNum("0", fmt.Sprintf("%d", length-1), seed)
	}

	nums[indexA] = a
	nums[indexB] = b

	strNums := generate_util.Arr2Str(nums)
	strResult := generate_util.Arr2Str([]int{indexA, indexB})
	//fmt.Printf("nums = %v\n", strNums)
	//fmt.Printf("target = %d\n", target)
	//fmt.Printf("indexA = %v\n", strResult)

	//templateStr := fmt.Sprintf(`{TestStr: "%v", CorrectResult: "@multiset%v",ShowWhenErr: " --by 系统生成"},`, fmt.Sprintf("%v\\n%d", strNums, target), strResult)
	templateStr := fmt.Sprintf(`  {
    "testStr": "%v",
    "correctResult": "@multiset%v",
    "showWhenErr": " --by 系统生成"
  }`, fmt.Sprintf("%v\\n%d", strNums, target), strResult)
	return templateStr
}
