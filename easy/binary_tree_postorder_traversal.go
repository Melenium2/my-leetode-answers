package easy

// Given the root of a binary tree, return the postorder traversal of its nodes' values.
//
//
//
// Example 1:
// Input: root = [1,null,2,3]
// Output: [3,2,1]
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
// Output: [2,1]
//
// Example 5:
// Input: root = [1,null,2]
// Output: [2,1]
func postorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)

	postorderTraversalUtil(root, &res)

	return res
}

func postorderTraversalUtil(root *TreeNode, sl *[]int) {
	if root == nil {
		return
	}

	postorderTraversalUtil(root.Left, sl)
	postorderTraversalUtil(root.Right, sl)

	*sl = append(*sl, root.Val)
}