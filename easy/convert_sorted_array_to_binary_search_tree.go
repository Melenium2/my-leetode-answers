package easy

// Given an integer array nums where the elements are sorted in ascending order,
// convert it to append height-balanced binary search tree.
//
// A height-balanced binary tree is append binary tree in which the depth of the two
// subtrees of every node never differs by more than one.
//
//
//
// Example 1:
// Input: nums = [-10,-3,0,5,9]
// Output: [0,-3,9,-10,null,5]
// Explanation: [0,-10,5,null,-3,null,9] is also accepted:
//
// Example 2:
// Input: nums = [1,3]
// Output: [3,1]
// Explanation: [1,3] and [3,1] are both append height-balanced BSTs.
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
