package algos

const inf = 1_000_000_000

type Node struct {
	Val    int
	Weight int
}

func Dijkstra(startFrom int, g [][]Node) []int {
	size := len(g)

	distances := initDistances(size)
	distances[startFrom-1] = 0

	visited := initVisited(size)

	for i := 0; i < size; i++ {
		vertex := -1

		for j := 0; j < size; j++ {
			if !visited[j] && (vertex == -1 || distances[j] < distances[vertex]) {
				vertex = j
			}
		}

		if distances[vertex] == inf {
			break
		}

		visited[vertex] = true

		edges := len(g[vertex])

		for j := 0; j < edges; j++ {
			to := g[vertex][j].Val - 1
			weight := g[vertex][j].Weight

			if to != -1 && distances[vertex]+weight < distances[to] {
				distances[to] = distances[vertex] + weight
			}
		}
	}

	return distances
}

func initDistances(size int) []int {
	dist := make([]int, size)

	for i := 0; i < size; i++ {
		dist[i] = inf
	}

	return dist
}

func initVisited(size int) []bool {
	return make([]bool, size)
}

func DefaultAdjacencyList() [][]Node {
	return [][]Node{
		{
			{Val: 2, Weight: 3}, {Val: 5, Weight: 5}, //.First Node
		},
		{
			{Val: 3, Weight: 1}, {Val: 4, Weight: 5}, // Second node
		},
		{
			{Val: 7, Weight: 3}, // Third node
		},
		{
			{Val: 3, Weight: 5}, {Val: 6, Weight: 3}, {Val: 7, Weight: 1}, // Forth node
		},
		{
			{Val: 4, Weight: 1}, {Val: 6, Weight: 7}, // Fifth node
		},
		{

			{},
		},
		{
			{},
		},
	}
}
