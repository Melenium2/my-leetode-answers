package medium

// You are given an integer array coins representing coins of different
// denominations and an integer amount representing a total amount of money.
//
// Return the number of combinations that make up that amount. If that amount
// of money cannot be made up by any combination of the coins, return 0.
//
// You may assume that you have an infinite number of each kind of coin.
//
// The answer is guaranteed to fit into a signed 32-bit integer.
//
//
//
// Example 1:
// Input: amount = 5, coins = [1,2,5]
// Output: 4
// Explanation: there are four ways to make up the amount:
// 5=5
// 5=2+2+1
// 5=2+1+1+1
// 5=1+1+1+1+1
//
// Example 2:
// Input: amount = 3, coins = [2]
// Output: 0
// Explanation: the amount of 3 cannot be made up just with coins of 2.
//
// Example 3:
// Input: amount = 10, coins = [10]
// Output: 1
func coinChangeTwo(amount int, coins []int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := 1; i < len(dp); i++ {
			if i >= coin {
				dp[i] += dp[i-coin]
			}
		}
	}

	return dp[len(dp)-1]
}

//				0  1  2  3  4  5  6  7
// [2, 3, 5] 2| 1  0  1  0  1  0  1  0
//           3| 1  0  1  1  1  1  2  1
//			 4| 1  0  1  1  1  2  2  2