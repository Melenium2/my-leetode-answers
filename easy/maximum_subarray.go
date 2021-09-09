package easy

// Given an integer array nums, find the contiguous subarray (containing at least one number)
// which has the largest sum and return its sum.
//
//
//
// Example 1:
// Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
// Output: 6
// Explanation: [4,-1,2,1] has the largest sum = 6.
//
// Example 2:
// Input: nums = [1]
// Output: 1
//
// Example 3:
// Input: nums = [5,4,-1,7,8]
// Output: 23
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		maxSum  = nums[0]
		currSum = maxSum
	)

	for i := 1; i < len(nums); i++ {
		currSum = max(currSum+nums[i], nums[i])
		maxSum = max(currSum, maxSum)
	}

	return maxSum
}

func max(i1, i2 int) int {
	if i1 > i2 {
		return i1
	}

	return i2
}
