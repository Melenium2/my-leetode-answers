package medium

type MaxHeap struct {
}

// Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.
//
//
//
// Example 1:
//
// Input: nums = [1,1,1,2,2,3], k = 2
// Output: [1,2]
//
// Example 2:
//
// Input: nums = [1], k = 1
// Output: [1]
//
//
// Constraints:
//
// 1 <= nums.length <= 105
// k is in the range [1, the number of unique elements in the array].
// It is guaranteed that the answer is unique.
//
//
// Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
func topKFrequent(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}

	freq := make([][]int, len(nums)+1)
	set := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		set[curr]++
	}

	for num, fr := range set {
		freq[fr] = append(freq[fr], num)
	}

	res := make([]int, 0)

	for i := len(freq) - 1; i >= 0; i-- {
		for j := 0; j < len(freq[i]); j++ {
			res = append(res, freq[i][j])
			if len(res) == k {
				return res
			}
		}
	}

	return nil
}
