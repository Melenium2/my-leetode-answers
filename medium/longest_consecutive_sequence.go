package medium

// Given an unsorted array of integers nums, return the length of the longest consecutive elements sequence.
//
// You must write an algorithm that runs in O(n) time.
//
//
//
// Example 1:
//
// Input: nums = [100,4,200,1,3,2]
// Output: 4
// Explanation: The longest consecutive elements sequence is [1, 2, 3, 4]. Therefore its length is 4.
//
// Example 2:
//
// Input: nums = [0,3,7,2,5,8,4,6,0,1]
// Output: 9
//
//
// Constraints:
//
// 0 <= nums.length <= 105
// -109 <= nums[i] <= 109
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{})
	longest := 0

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		set[curr] = struct{}{}
	}

	for i := 0; i < len(nums); i++ {
		curr := nums[i]

		_, ok := set[curr-1]
		if ok {
			continue
		}

		currLength := 1

		for {
			_, ok := set[curr+currLength]
			if !ok {
				break
			}

			currLength++
		}

		longest = max(longest, currLength)
	}

	return longest
}
