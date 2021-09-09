package easy

// You have n coins and you want to build a staircase with these coins.
// The staircase consists of k rows where the ith row has exactly i coins.
// The last row of the staircase may be incomplete.
//
// Given the integer n, return the number of complete rows of the staircase you will build.
//
//
//
// Example 1:
// Input: n = 5
// Output: 2
// Explanation: Because the 3rd row is incomplete, we return 2.
//
// Example 2:
// Input: n = 8
// Output: 3
// Explanation: Because the 4th row is incomplete, we return 3.
func arrangeCoins(n int) int {
	l, r := 0, n

	for l <= r {
		mid := l + (r-l)/2
		curr := mid * (mid+1) / 2

		if curr == n {
			return mid
		}

		if n < curr {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return r
}
