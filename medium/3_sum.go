package medium

import "sort"

// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]]
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Notice that the solution set must not contain duplicate triplets.
//
//
//
// Example 1:
//
// Input: nums = [-1,0,1,2,-1,-4]
// Output: [[-1,-1,2],[-1,0,1]]
//
// Example 2:
//
// Input: nums = []
// Output: []
//
// Example 3:
//
// Input: nums = [0]
// Output: []
//
//
// Constraints:
//
// 0 <= nums.length <= 3000
// -105 <= nums[i] <= 105
func threeSum(nums []int) [][]int {
	if len(nums) == 0 {
		return nil
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := make([][]int, 0)

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		if i > 0 && curr == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1

		for l < r {
			sum := curr + nums[l] + nums[r]

			if sum > 0 {
				r--

				continue
			}

			if sum < 0 {
				l++

				continue
			}

			res = append(res, []int{curr, nums[l], nums[r]})
			l++

			for nums[l] == nums[l-1] && l < r {
				l++
			}
		}
	}

	return res
}
