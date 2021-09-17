package easy

import "math"

// Given the root of a binary search tree (BST) with duplicates, return
// all the mode(s) (i.e., the most frequently occurred element) in it.
//
// If the tree has more than one mode, return them in any order.
//
// Assume a BST is defined as follows:
//
// The left subtree of a node contains only nodes with keys less than or equal to the node's key.
// The right subtree of a node contains only nodes with keys greater than or equal to the node's key.
// Both the left and right subtrees must also be binary search trees.
//
//
// Example 1:
//
//
// Input: root = [1,null,2,2]
// Output: [2]
// Example 2:
//
// Input: root = [0]
// Output: [0]
//
//
// Constraints:
//
// The number of nodes in the tree is in the range [1, 104].
// -105 <= Node.val <= 105
//
//
// Follow up: Could you do that without using any extra space? (Assume that the
// implicit stack space incurred due to recursion does not count).
func findMode(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	elements := make(map[int]int)

	findModeTraversal(root, elements)

	res := make([]int, 0, len(elements))

	maxMode := math.MinInt
	for _, v := range elements {
		maxMode = max(maxMode, v)
	}
	for k, v := range elements {
		if v == maxMode {
			res = append(res, k)
		}
	}

	return res
}

func findModeTraversal(root *TreeNode, elements map[int]int) {
	if root == nil {
		return
	}

	findModeTraversal(root.Left, elements)
	elements[root.Val]++
	findModeTraversal(root.Right, elements)
}
