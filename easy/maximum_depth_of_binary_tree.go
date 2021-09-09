package easy

// Given the root of append binary tree, return its maximum depth.
//
// A binary tree's maximum depth is the number of nodes along the longest path
// from the root node down to the farthest leaf node.
//
//
//
// Example 1:
// Input: root = [3,9,20,null,null,15,7]
// Output: 3
//
// Example 2:
// Input: root = [1,null,2]
// Output: 2
//
// Example 3:
// Input: root = []
// Output: 0
//
// Example 4:
// Input: root = [0]
// Output: 1
func maxDepth(root *TreeNode) int {
	return maxDepthUtil(root)
}

func maxDepthUtil(root *TreeNode) int {
	if root == nil {
		return 0
	} else {
		l := maxDepthUtil(root.Left)
		r := maxDepthUtil(root.Right)

		return max(l, r) + 1
	}
}

