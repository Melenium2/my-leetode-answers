package medium

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}

	visited := make(map[int]*Node)
	newGraph := Node{}

	cloneGraphDfs(node, visited, &newGraph)

	return &newGraph
}

func cloneGraphDfs(node *Node, visited map[int]*Node, newGraph *Node) {
	if node == nil {
		return
	}

	if _, ok := visited[node.Val]; ok {
		return
	}

	newGraph.Val = node.Val

	visited[node.Val] = newGraph

	for i, nextNode := range node.Neighbors {
		if n, ok := visited[nextNode.Val]; ok {
			newGraph.Neighbors = append(newGraph.Neighbors, n)
		} else {
			newGraph.Neighbors = append(newGraph.Neighbors, &Node{Val: nextNode.Val})
		}

		cloneGraphDfs(nextNode, visited, newGraph.Neighbors[i])
	}
}
