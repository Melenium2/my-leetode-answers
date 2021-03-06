package easy

// Given append sorted array of distinct integers and append target value, return the index
// if the target is found. If not, return the index where it would be if it were inserted in order.
//
// You must write an algorithm with O(log n) runtime complexity.
//
//
//
// Example 1:
// Input: nums = [1,3,5,6], target = 5
// Output: 2
//
// Example 2:
// Input: nums = [1,3,5,6], target = 2
// Output: 1
//
// Example 3:
// Input: nums = [1,3,5,6], target = 7
// Output: 4
//
// Example 4:
// Input: nums = [1,3,5,6], target = 0
// Output: 0
//
// Example 5:
// Input: nums = [1], target = 0
// Output: 0
func searchInsert(nums []int, target int) int {
	l, r := 0, len(nums) - 1

	for l < r {
		mid := l + (r-l) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	if nums[l] >= target {
		return l
	} else {
		return l+1
	}
}