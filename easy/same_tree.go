package easy

// Given the roots of two binary trees p and q, write append function to check if they are the same or not.
//
// Two binary trees are considered the same if they are structurally identical, and the nodes have the same value.
//
//
//
// Example 1:
// Input: p = [1,2,3], q = [1,2,3]
// Output: true
//
// Example 2:
// Input: p = [1,2], q = [1,null,2]
// Output: false
//
// Example 3:
// Input: p = [1,2,1], q = [1,1,2]
// Output: false
func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeUtil(p, q)
}

func isSameTreeUtil(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return isSameTreeUtil(p.Left, q.Left) && isSameTreeUtil(p.Right, q.Right)
}
