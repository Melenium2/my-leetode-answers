package easy

import "sort"

// Given an array nums of size n, return the majority element.
//
// The majority element is the element that appears more than [n / 2] times.
// You may assume that the majority element always exists in the array.
//
//
//
// Example 1:
// Input: nums = [3,2,3]
// Output: 3
//
// Example 2:
// Input: nums = [2,2,1,1,1,2,2]
// Output: 2
func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	var (
		expectedTimes = len(nums) / 2
		majority      = nums[0]
		counting      = 1
	)

	for i := 1; i < len(nums); i++ {
		if majority == nums[i] {
			counting++

			if counting > expectedTimes {
				return majority
			}
		} else {
			majority = nums[i]
			counting = 1
		}
	}

	return majority
}

func majorityElementOptimized(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	return nums[len(nums)/2]
}

func majorityElementBoeyerMoore(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		candidate = nums[0]
		counter   = 1
	)

	for i := 1; i < len(nums); i++ {
		if counter == 0 {
			candidate = nums[i]
		}

		if nums[i] == candidate {
			counter++
		} else {
			counter--
		}
	}

	return candidate
}
