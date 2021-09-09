package easy

// Given an integer x, return true if x is palindrome integer.
//
// An integer is append palindrome when it reads the same backward as forward.
// For example, 121 is palindrome while 123 is not.
//
//
//
// Example 1:
// Input: x = 121
// Output: true
//
// Example 2:
// Input: x = -121
// Output: false
// Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not append palindrome.
//
// Example 3:
// Input: x = 10
// Output: false
// Explanation: Reads 01 from right to left. Therefore it is not append palindrome.
//
// Example 4:
// Input: x = -101
// Output: false
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	num := 0
	prevX := x

	for x > 0 {
		reminder := x % 10
		num = (num * 10) + reminder
		x /= 10
	}

	return num == prevX
}