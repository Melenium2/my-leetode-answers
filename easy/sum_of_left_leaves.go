package easy

// Given the root of a binary tree, return the sum of all left leaves.
//
//
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 24
// Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.
//
// Example 2:
// Input: root = [1]
// Output: 0
func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0

	if root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			sum += root.Left.Val
		} else {
			sum += sumOfLeftLeaves(root.Left)
		}
	}

	if root.Right != nil {
		if root.Right.Left != nil || root.Right.Right != nil {
			sum += sumOfLeftLeaves(root.Right)
		}
	}

	return sum
}