package easy

// Given the roots of two binary trees root and subRoot, return true if there
// is a subtree of root with the same structure and node values of subRoot and false otherwise.
//
// A subtree of a binary tree tree is a tree that consists of a node in tree and all of
// this node's descendants. The tree tree could also be considered as a subtree of itself.
//
//
//
// Example 1:
// Input: root = [3,4,5,1,2], subRoot = [4,1,2]
// Output: true
//
// Example 2:
// Input: root = [3,4,5,1,2,null,null,null,null,0], subRoot = [4,1,2]
// Output: false
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil || subRoot == nil {
		return false
	}

	return isSubtreeUtill(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSubtreeUtill(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val != subRoot.Val {
		return false
	}

	return isSubtreeUtill(root.Left, subRoot.Left) && isSubtreeUtill(root.Right, subRoot.Right)
}
