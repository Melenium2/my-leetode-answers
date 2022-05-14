package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var (
		res = make([][]int, 0)
		q   = make([]*TreeNode, 0)
	)

	q = append(q, root)

	for len(q) > 0 {
		count := len(q)
		nodesAtLevel := make([]int, 0, count)

		for i := 0; i < count; i++ {
			curr := q[0]
			q = q[1:]

			nodesAtLevel = append(nodesAtLevel, curr.Val)

			if curr.Left != nil {
				q = append(q, curr.Left)
			}

			if curr.Right != nil {
				q = append(q, curr.Right)
			}
		}

		res = append(res, nodesAtLevel)
	}

	return res
}
