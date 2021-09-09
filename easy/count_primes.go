package easy

// Count the number of prime numbers less than a non-negative number, n.
//
//
//
// Example 1:
//
// Input: n = 10
// Output: 4
// Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
// Example 2:
//
// Input: n = 0
// Output: 0
// Example 3:
//
// Input: n = 1
// Output: 0
func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	bools := make([]bool, n)

	for i := 2; i*i < n; i++ {
		if !bools[i] {
			for j := i*i; j < n; j+=i {
				bools[j] = true
			}
		}
	}

	count := 0
	for i := 0; i < n; i++ {
		if !bools[i] {
			count++
		}
	}

	return count-2
}