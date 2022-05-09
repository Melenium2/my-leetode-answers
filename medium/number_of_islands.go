package medium

type position struct {
	x int
	y int
}

// Given an m x n 2D binary grid grid which represents a map of '1's
//(land) and '0's (water), return the number of islands.
//
// An island is surrounded by water and is formed by connecting adjacent
//lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.
//
//
//
// Example 1:
//
// Input: grid = [
//  ["1","1","1","1","0"],
//  ["1","1","0","1","0"],
//  ["1","1","0","0","0"],
//  ["0","0","0","0","0"]
// ]
// Output: 1
// Example 2:
//
// Input: grid = [
//  ["1","1","0","0","0"],
//  ["1","1","0","0","0"],
//  ["0","0","1","0","0"],
//  ["0","0","0","1","1"]
// ]
// Output: 3
//
//
// Constraints:
//
// m == grid.length
// n == grid[i].length
// 1 <= m, n <= 300
// grid[i][j] is '0' or '1'.
func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	visited := make(map[position]struct{})
	islands := 0

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c]-'0' == 0 {
				continue
			}

			if _, ok := visited[position{r, c}]; ok {
				continue
			}

			findIsland(r, c, grid, visited)
			islands++
		}
	}

	return islands
}

func findIsland(r, c int, grid [][]byte, visited map[position]struct{}) {
	rows, cols := len(grid), len(grid[0])

	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	q := make([][]int, 0)

	visited[position{r, c}] = struct{}{}

	q = append(q, []int{r, c})

	for len(q) > 0 {
		row, coll := q[0][0], q[0][1]
		q = append(q[:0], q[1:]...)

		for _, direction := range directions {
			rr, cc := row+direction[0], coll+direction[1]

			_, alreadyVisited := visited[position{rr, cc}]

			if rr >= 0 && rr < rows && cc >= 0 && cc < cols && grid[rr][cc]-'0' == 1 && !alreadyVisited {
				q = append(q, []int{rr, cc})
				visited[position{rr, cc}] = struct{}{}
			}
		}
	}
}
