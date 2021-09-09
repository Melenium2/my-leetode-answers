package easy

// Given a positive integer num, write a function which returns True if num is a perfect square else False.
//
// Follow up: Do not use any built-in library function such as sqrt.
//
//
//
// Example 1:
//
// Input: num = 16
// Output: true
// Example 2:
//
// Input: num = 14
// Output: false
func isPerfectSquare(num int) bool {
	if num < 2 {
		return false
	}

	l, r := 0, num/2

	for l <= r {
		mid := l + (r-l)/2

		sqr := mid * mid

		if sqr == num {
			return true
		}

		if sqr < num {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return false
}
