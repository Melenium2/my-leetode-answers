package easy

// Given an integer rowIndex, return the rowIndexth (0-indexed) row of the Pascal's triangle.
//
// In Pascal's triangle, each number is the sum of the two numbers directly above it as shown:
//
//
//
//
// Example 1:
//
// Input: rowIndex = 3
// Output: [1,3,3,1]
// Example 2:
//
// Input: rowIndex = 0
// Output: [1]
// Example 3:
//
// Input: rowIndex = 1
// Output: [1,1]
func getPascalTriangleRow(rowIndex int) []int {
	row := make([]int, rowIndex+1)

	row[0] = 1

	if rowIndex == 0 {
		return row
	}

	for i := 1; i < rowIndex/2+1; i++ {
		row[i] = row[i-1] * (rowIndex - i + 1) / i
	}

	for i := rowIndex/2 + 1; i < rowIndex; i++ {
		row[i] = row[rowIndex-i]
	}

	row[len(row)-1] = 1

	return row
}
