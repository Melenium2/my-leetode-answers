package medium

import "math"

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isValidBSTUtil(root, math.MinInt, math.MaxInt)
}

func isValidBSTUtil(root *TreeNode, left, right int) bool {
	if root == nil {
		return true
	}

	if !(root.Val < right && root.Val > left) {
		return false
	}

	return isValidBSTUtil(root.Left, left, root.Val) && isValidBSTUtil(root.Right, root.Val, right)
}
