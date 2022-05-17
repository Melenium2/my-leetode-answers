package medium

func kthSmallest(root *TreeNode, k int) int {
	nums := make([]int, 0)

	kthSmallestInOrder(root, &nums)

	return nums[k-1]
}

func kthSmallestInOrder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}

	kthSmallestInOrder(root.Left, nums)
	*nums = append(*nums, root.Val)
	kthSmallestInOrder(root.Right, nums)
}
