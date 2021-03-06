package medium

// You are given an integer array coins representing coins of different
// denominations and an integer amount representing a total amount of money.
//
// Return the fewest number of coins that you need to make up that amount.
// If that amount of money cannot be made up by any combination of the coins, return -1.
//
// You may assume that you have an infinite number of each kind of coin.
//
//
//
// Example 1:
// Input: coins = [1,2,5], amount = 11
// Output: 3
// Explanation: 11 = 5 + 5 + 1
//
// Example 2:
// Input: coins = [2], amount = 3
// Output: -1
//
// Example 3:
// Input: coins = [1], amount = 0
// Output: 0
//
// Example 4:
// Input: coins = [1], amount = 1
// Output: 1
//
// Example 5:
// Input: coins = [1], amount = 2
// Output: 2
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}

	dp[0] = 0

	for i := 1; i < len(dp); i++ {
		for j := 0; j < len(coins); j++ {
			coin := coins[j]

			if i-coin >= 0 {
				dp[i] = min(dp[i], 1+dp[i-coin])
			}
		}
	}

	if dp[amount] != amount+1 {
		return dp[amount]
	} else {
		return -1
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
