package easy

// Given the root of a binary tree, invert the tree, and return its root.
//
//
//
// Example 1:
//
//
// Input: root = [4,2,7,1,3,6,9]
// Output: [4,7,2,9,6,3,1]
// Example 2:
//
//
// Input: root = [2,1,3]
// Output: [2,3,1]
// Example 3:
//
// Input: root = []
// Output: []
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return root
	}

	root.Left, root.Right = root.Right, root.Left

	root.Left = invertTree(root.Left)
	root.Right = invertTree(root.Right)

	return root
}