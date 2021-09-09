package easy

// Given two integer arrays nums1 and nums2, return an array of their intersection.
// Each element in the result must be unique and you may return the result in any order.
//
//
//
// Example 1:
// Input: nums1 = [1,2,2,1], nums2 = [2,2]
// Output: [2]
//
// Example 2:
// Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
// Output: [9,4]
// Explanation: [4,9] is also accepted.
func intersection(nums1 []int, nums2 []int) []int {
	set1 := make(map[int]struct{})

	for i := 0; i < len(nums1); i++ {
		set1[nums1[i]] = struct{}{}
	}

	intersections := make([]int, 0, len(nums2))

	for i := 0; i < len(nums2); i++ {
		if _, ok := set1[nums2[i]]; ok {
			intersections = append(intersections, nums2[i])
			delete(set1, nums2[i])
		}
	}

	return intersections
}