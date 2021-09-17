package easy

import "fmt"

// Given an integer num, return a string of its base 7 representation.
//
//
//
// Example 1:
// Input: num = 100
// Output: "202"
//
// Example 2:
// Input: num = -7
// Output: "-10"
func convertToBase7(num int) string {
	sign := 1
	if num < 0 {
		sign = -1
		num *= sign
	}

	res := ""
	for num >= 7 {
		res = fmt.Sprintf("%d", num%7) + res
		num /= 7
	}

	res = fmt.Sprintf("%d", num) + res

	if sign > 0 {
		return res
	} else {
		return fmt.Sprintf("-%s", res)
	}
}
