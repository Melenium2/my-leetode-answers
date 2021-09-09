package easy

// Given the root of a binary tree, return the preorder traversal of its nodes' values.
//
//
//
// Example 1:
// Input: root = [1,null,2,3]
// Output: [1,2,3]
//
// Example 2:
// Input: root = []
// Output: []
//
// Example 3:
// Input: root = [1]
// Output: [1]
//
// Example 4:
// Input: root = [1,2]
// Output: [1,2]
//
// Example 5:
// Input: root = [1,null,2]
// Output: [1,2]
func preorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	preorderTraversalUtil(root, &res)

	return res
}

func preorderTraversalUtil(root *TreeNode, sl *[]int) {
	if root == nil {
		return
	}

	*sl = append(*sl, root.Val)

	preorderTraversalUtil(root.Left, sl)
	preorderTraversalUtil(root.Right, sl)
}