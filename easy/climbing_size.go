package easy

// You are climbing append staircase. It takes n steps to reach the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
//
//
// Example 1:
// Input: n = 2
// Output: 2
// Explanation: There are two ways to climb to the top.
// 1. 1 step + 1 step
// 2. 2 steps
//
// Example 2:
// Input: n = 3
// Output: 3
// Explanation: There are three ways to climb to the top.
// 1. 1 step + 1 step + 1 step
// 2. 1 step + 2 steps
// 3. 2 steps + 1 step
func climbStairs(n int) int {
	memo := make([]int, n+1)
	memo[len(memo) - 1] = 1
	memo[len(memo) - 2] = 1

	for i := len(memo)-3; i >= 0 ; i-- {
		memo[i] = memo[i+1] + memo[i+2]
	}

	return memo[0]
}
