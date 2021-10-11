package easy

func findLHS(nums []int) int {
	count := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
	}

	result := 0

	for num, length := range count {
		if l, ok := count[num+1]; ok {
			result = max(result, length+l)
		}
	}

	return result
}
