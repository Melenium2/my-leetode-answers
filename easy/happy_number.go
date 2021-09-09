package easy

// Write an algorithm to determine if a number n is happy.
//
// A happy number is a number defined by the following process:
//
// Starting with any positive integer, replace the number by the sum of the squares of its digits.
// Repeat the process until the number equals 1 (where it will stay),
// or it loops endlessly in a cycle which does not include 1.
// Those numbers for which this process ends in 1 are happy.
// Return true if n is a happy number, and false if not.
//
//
//
// Example 1:
// Input: n = 19
// Output: true
// Explanation:
// 1^2 + 9^2 = 82
// 8^2 + 2^2 = 68
// 6^2 + 8^2 = 100
// 1^2 + 0^2 + 0^2 = 1
//
// Example 2:
// Input: n = 2
// Output: false
func isHappy(n int) bool {
	var (
		resMemo  = make(map[int]struct{})
		sumOfPow = n
	)

	for {
		nums := separateNumbers(sumOfPow)
		sumOfPow = sumPow(nums)

		if sumOfPow == 1 {
			break
		}

		if _, ok := resMemo[sumOfPow]; ok {
			return false
		}

		resMemo[sumOfPow] = struct{}{}
	}

	return true
}

func fastPow(x, y int) int {
	res := 1

	for y != 0 {
		if y&1 == 1 {
			res *= x
		}

		y >>= 1

		x = x * x
	}

	return res
}

func separateNumbers(num int) []int {
	ints := make([]int, 0)

	for num != 0 {
		ints = append(ints, num%10)
		num /= 10
	}

	return ints
}

func sumPow(numbs []int) int {
	sum := 0

	for i := 0; i < len(numbs); i++ {
		sum += fastPow(numbs[i], 2)
	}

	return sum
}
