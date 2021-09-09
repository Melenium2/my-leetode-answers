package easy

// Given an integer numRows, return the first numRows of Pascal's triangle.
//
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
//
//
//
//
// Example 1:
// Input: numRows = 5
// Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
//
// Example 2:
// Input: numRows = 1
// Output: [[1]]
func pascalTriangle(numRows int) [][]int {
	if numRows == 0 {
		return nil
	}

	triangle := make([][]int, numRows)

	triangle[0] = []int{1}

	if numRows == 1 {
		return triangle
	}

	for level := 1; level < numRows; level++ {
		levelSlice := make([]int, level+1)

		for i := 1; i < level; i++ {
			levelSlice[i] = triangle[level-1][i-1] + triangle[level-1][i]
		}

		levelSlice[0], levelSlice[len(levelSlice)-1] = 1, 1
		triangle[level] = levelSlice
	}

	return triangle
}
