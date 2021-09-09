package easy

// You are given row x col grid representing a map where grid[i][j] = 1
// represents land and grid[i][j] = 0 represents water.
//
// Grid cells are connected horizontally/vertically (not diagonally).
// The grid is completely surrounded by water, and there is exactly one island (i.e., one or more connected land cells).
//
// The island doesn't have "lakes", meaning the water inside isn't connected
// to the water around the island. One cell is a square with side length 1. The grid is rectangular, width and height don't exceed 100. Determine the perimeter of the island.
//
//
//
// Example 1:
// Input: grid = [[0,1,0,0],[1,1,1,0],[0,1,0,0],[1,1,0,0]]
// Output: 16
// Explanation: The perimeter is the 16 yellow stripes in the image above.
//
// Example 2:
// Input: grid = [[1]]
// Output: 4
//
// Example 3:
// Input: grid = [[1,0]]
// Output: 4
func islandPerimeter(grid [][]int) int {
	visited := make(map[[2]int]struct{})

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				return dfs(visited, i, j, grid)
			}
		}
	}

	return 0
}

func dfs(visited map[[2]int]struct{}, i, j int, grid [][]int) int {
	if i >= len(grid) || j >= len(grid[0]) || i < 0 || j < 0 || grid[i][j] == 0 {
		return 1
	}
	if _, ok := visited[[2]int{i, j}]; ok {
		return 0
	}

	visited[[2]int{i, j}] = struct{}{}

	perim := dfs(visited, i, j+1, grid)
	perim += dfs(visited, i, j-1, grid)
	perim += dfs(visited, i + 1, j, grid)
	perim += dfs(visited, i-1, j, grid)

	return perim
}
