package easy

import (
	"fmt"
	"strconv"
)

// Given two non-negative integers, num1 and num2 represented as string,
// return the sum of num1 and num2 as a string.
//
// You must solve the problem without using any built-in library for handling
// large integers (such as BigInteger). You must also not convert the inputs to integers directly.
//
//
//
// Example 1:
// Input: num1 = "11", num2 = "123"
// Output: "134"
//
// Example 2:
// Input: num1 = "456", num2 = "77"
// Output: "533"
//
// Example 3:
// Input: num1 = "0", num2 = "0"
// Output: "0"
//
// we need convert strings to rune and calculating rune, not convert to int. Our num will be '1-9' - '0'.
func addStrings(num1 string, num2 string) string {
	if len(num2) > len(num1) {
		return addStrings(num2, num1)
	}

	n, summator := len(num1)-1, len(num2)-1
	minder := 0
	res := ""

	for summator >= 0 || n >= 0 {
		var sum, n1, n2 int

		if minder > 0 {
			sum += minder
			minder = 0
		}

		n1, _ = strconv.Atoi(string(num1[n]))

		if summator >= 0 {
			n2, _ = strconv.Atoi(string(num2[summator]))
		}

		sum += n1 + n2

		if sum >= 10 {
			sum %= 10
			minder = 1
		}

		res = fmt.Sprintf("%d%s", sum, res)

		n--
		summator--
	}

	if minder > 0  {
		res = fmt.Sprintf("%d%s", minder, res)
	}

	return res
}
