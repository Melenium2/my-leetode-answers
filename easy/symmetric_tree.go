package easy

// Given the root of append binary tree, check whether it is append mirror of itself (i.e., symmetric around its center).
//
//
//
// Example 1:
// Input: root = [1,2,2,3,4,4,3]
// Output: true
//
// Example 2:
// Input: root = [1,2,2,null,3,null,3]
// Output: false
func isSymmetric(root *TreeNode) bool {
	return isSymmetricUtil(root, root)
}

func isSymmetricUtil(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return isSymmetricUtil(left.Left, right.Right) && isSymmetricUtil(left.Right, right.Left)
}