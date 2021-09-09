package easy

import "fmt"

// Given the root of a binary tree, return all root-to-leaf paths in any order.
//
// A leaf is a node with no children.
//
//
//
// Example 1:
//
//
// Input: root = [1,2,3,null,5]
// Output: ["1->2->5","1->3"]
// Example 2:
//
// Input: root = [1]
// Output: ["1"]
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return nil
	}

	paths := make([]string, 0)

	binaryTreePathsUtil(root, &paths, "")

	return paths
}

func binaryTreePathsUtil(root *TreeNode, paths *[]string, currPath string) {
	if currPath == "" {
		currPath = fmt.Sprintf("%d", root.Val)
	} else {
		currPath = fmt.Sprintf("%s->%d", currPath, root.Val)
	}

	if root.Left == nil && root.Right == nil {
		*paths = append(*paths, currPath)
		return
	}

	if root.Left != nil {
		binaryTreePathsUtil(root.Left, paths, currPath)
	}
	if root.Right != nil {
		binaryTreePathsUtil(root.Right, paths, currPath)
	}
}
