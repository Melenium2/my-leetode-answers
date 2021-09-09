package easy

// Given append non-negative integer x, compute and return the square root of x.
//
// Since the return type is an integer, the decimal digits are truncated, and only the integer part of the result is returned.
//
// Note: You are not allowed to use any built-in exponent function or operator, such as pow(x, 0.5) or x ** 0.5.
//
//
//
// Example 1:
// Input: x = 4
// Output: 2
//
// Example 2:
// Input: x = 8
// Output: 2
// Explanation: The square root of 8 is 2.82842..., and since the decimal part is truncated, 2 is returned.
func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	l, r := 1, x

	for l < r {
		mid := l + (r-l) / 2

		sqr := mid * mid

		if sqr == x {
			return mid
		}

		if sqr > x {
			r = mid
		} else if sqr < x {
			l = mid + 1
		}
	}

	return l - 1
}