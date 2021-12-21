package medium

// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. You are given an array
// prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi first if you want to take
// course ai.
//
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.
//
//
//
// Example 1:
//
// Input: numCourses = 2, prerequisites = [[1,0]]
// Output: true
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0. So it is possible.
//
// Example 2:
//
// Input: numCourses = 2, prerequisites = [[1,0],[0,1]]
// Output: false
// Explanation: There are a total of 2 courses to take.
// To take course 1 you should have finished course 0, and to take course 0 you should also have finished course 1.
// So it is impossible.
//
//
// Constraints:
//
// 1 <= numCourses <= 105
// 0 <= prerequisites.length <= 5000
// prerequisites[i].length == 2
// 0 <= ai, bi < numCourses
// All the pairs prerequisites[i] are unique.
func canFinish(numCourses int, prerequisites [][]int) bool {
	adjMap := make(map[int][]int, numCourses)

	for _, prerequisite := range prerequisites {
		from, to := prerequisite[0], prerequisite[1]

		adjMap[from] = append(adjMap[from], to)
	}

	visited := make(map[int]struct{})

	for i := 0; i < numCourses; i++ {
		if adjMap[i] == nil {
			continue
		}

		if !adjacencyDFS(adjMap, i, visited) {
			return false
		}
	}

	return true
}

func adjacencyDFS(prerequisite map[int][]int, element int, visited map[int]struct{}) bool {
	if _, ok := visited[element]; ok {
		return false
	}

	if len(prerequisite[element]) == 0 {
		return true
	}

	visited[element] = struct{}{}

	for i := 0; i < len(prerequisite[element]); i++ {
		if !adjacencyDFS(prerequisite, prerequisite[element][i], visited) {
			return false
		}
	}

	delete(visited, element)
	prerequisite[element] = []int{}

	return true
}
