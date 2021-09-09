package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Given the root of append binary tree, return the inorder traversal of its nodes' values.
//
// Example 1:
// Input: root = [1,null,2,3]
// Output: [1,3,2]
//
// Example 2:
// Input: root = []
// Output: []
//
// Example 3:
// Input: root = [1]
// Output: [1]
func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	traversal(&res, root)

	return res
}

func traversal(values *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	traversal(values, root.Left)

	*values = append(*values, root.Val)

	traversal(values, root.Right)
}
