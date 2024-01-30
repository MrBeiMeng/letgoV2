package IDzzzz_two_sum

func twoSum(nums []int, target int) []int {
	valueMap := make(map[int]int)
	for i, num := range nums {
		if index, ok := valueMap[target-num]; ok {
			return []int{i, index}
		}
		valueMap[num] = i
	}

	return nil
}
