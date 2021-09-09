package easy

import "math"

// Given a string columnTitle that represents the column title as appear in an Excel sheet,
// return its corresponding column number.
//
// For example:
//
// A -> 1
// B -> 2
// C -> 3
// ...
// Z -> 26
// AA -> 27
// AB -> 28
// ...
//
//
// Example 1:
//
// Input: columnTitle = "A"
// Output: 1
// Example 2:
//
// Input: columnTitle = "AB"
// Output: 28
// Example 3:
//
// Input: columnTitle = "ZY"
// Output: 701
// Example 4:
//
// Input: columnTitle = "FXSHRXW"
// Output: 2147483647
func titleToNumber(columnTitle string) int {
	res := 0

	count := 0
	for i := len(columnTitle) - 1; i >= 0; i-- {
		res += int(math.Pow(26, float64(count))) * int((columnTitle[i] - 'A' + 1))
		count++
	}

	return res
}
