package IDzzzz_two_sum

import (
	"fmt"
	"letgoV2/system_code/pkg/func_operator"
	"letgoV2/system_code/pkg/logging"
	"testing"
)

func Test_twoSum(t *testing.T) {

	for i, sampleTest := range sampleTests { // 你可以从meta_data.go 中找到测试集合
		t.Run(fmt.Sprintf("CASE %d", i+1), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("发生了panic:", r)
					logging.Error(r)
					t.Error(r)
				}
			}()

			err, result := func_operator.RunFunc(twoSum, sampleTest)
			if err != nil {
				logging.Error(err)
				t.Error(err.Error())
			}

			logging.Info(fmt.Sprintf("longestPalindrome(%s) = %v ", sampleTest, result))
		})
	}

}
