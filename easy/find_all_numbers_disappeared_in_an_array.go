package easy

// Given an array nums of n integers where nums[i] is in the range [1, n],
// return an array of all the integers in the range [1, n] that do not appear in nums.
//
//
//
// Example 1:
// Input: nums = [4,3,2,7,8,2,3,1]
// Output: [5,6]
//
// Example 2:
// Input: nums = [1,1]
// Output: [2]
//
//
// Constraints:
//
// n == nums.length
// 1 <= n <= 10^5
// 1 <= nums[i] <= n
//
//
// Follow up: Could you do it without extra space and in O(n) runtime? You may assume
// the returned list does not count as extra space.
func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		for {
			index := nums[i] - 1
			if nums[index] == nums[i] {
				break
			}
			nums[i], nums[index] = nums[index], nums[i]
		}
	}

	answer := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		if nums[i] == i+1 {
			answer = append(answer, i+1)
		}
	}

	return answer
}
