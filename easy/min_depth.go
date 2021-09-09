package easy

import "math"

// Given a binary tree, find its minimum depth.
//
// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.
//
// Note: A leaf is a node with no children.
//
//
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 2
//
// Example 2:
// Input: root = [2,null,3,null,4,null,5,null,6]
// Output: 5
func minDepth(root *TreeNode) int {
	return minDepthUtil(root)
}

func minDepthUtil(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil && root.Right == nil {
		return 1
	}

	var l, r int = math.MaxInt32, math.MaxInt32

	if root.Left != nil {
		l = minDepthUtil(root.Left)
	}

	if root.Right != nil {
		r = minDepthUtil(root.Right)
	}

	return min(l, r) + 1
}

func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}

	return v2
}
