package easy

import "math"

// Given the root of a Binary Search Tree (BST), return the minimum absolute
// difference between the values of any two different nodes in the tree.
//
//
//
// Example 1:
// Input: root = [4,2,6,1,3]
// Output: 1
//
// Example 2:
// Input: root = [1,0,48,null,null,12,49]
// Output: 1
func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}

	min := math.MaxInt32
	prev := -1_000_000

	minimumDifferenceTraverse(root, &min, &prev)

	return min
}

func minimumDifferenceTraverse(root *TreeNode, minimum *int, prev *int) {
	if root == nil {
		return
	}

	minimumDifferenceTraverse(root.Left, minimum, prev)

	*minimum = min(*minimum, root.Val-*prev)
	*prev = root.Val

	minimumDifferenceTraverse(root.Right, minimum, prev)
}
