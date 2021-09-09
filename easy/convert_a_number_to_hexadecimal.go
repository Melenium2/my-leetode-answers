package easy

import (
	"fmt"
	"math"
)

// Given an integer num, return a string representing its hexadecimal representation.
// For negative integers, two’s complement method is used.
//
// All the letters in the answer string should be lowercase characters, and there should
// not be any leading zeros in the answer except for the zero itself.
//
// Note: You are not allowed to use any built-in library method to directly solve this problem.
//
//
//
// Example 1:
// Input: num = 26
// Output: "1a"
//
// Example 2:
// Input: num = -1
// Output: "ffffffff"
func toHex(num int) string {
	if num < 0 {
		num = int(math.Pow(2, 32)) - (-1 * num)
	}

	hex := "0123456789abcdef"

	remainder := 0
	remainderStr := ""

	for {
		remainder = num % 16
		num /= 16

		remainderStr = string(hex[remainder]) + remainderStr

		if num < 15 {
			break
		}
	}

	if num != 0 {
		return fmt.Sprintf("%s%s", string(hex[num]), remainderStr)
	} else {
		return remainderStr
	}
}
