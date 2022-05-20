package medium

// You are given a string s and an integer k. You can choose any character of the string and change it
// to any other uppercase English character. You can perform this operation at most k times.
//
// Return the length of the longest substring containing the same letter you can get after performing the
// above operations.
//
//
//
// Example 1:
//
// Input: s = "ABAB", k = 2
// Output: 4
// Explanation: Replace the two 'A's with two 'B's or vice versa.
//
// Example 2:
//
// Input: s = "AABABBA", k = 1
// Output: 4
// Explanation: Replace the one 'A' in the middle with 'B' and form "AABBBBA".
// The substring "BBBB" has the longest repeating letters, which is 4.
//
//
// Constraints:
//
// 1 <= s.length <= 105
// s consists of only uppercase English letters.
// 0 <= k <= s.length
func characterReplacement(s string, k int) int {
	if len(s) == 0 {
		return 0
	}

	// ABAB k = 2
	// AABABBA k = 1

	var (
		charSet [26]int
		maxFreq = 0
		start   = 0
		maxLen  = 0
	)

	for end := 0; end < len(s); end++ {
		curr := int(s[end] - 'A')
		charSet[curr]++

		maxFreq = max(maxFreq, charSet[curr])

		if end-start+1 > maxFreq+k {
			charSet[s[start]-'A']--
			start++
		}

		maxLen = max(maxLen, end-start+1)
	}

	return maxLen
}
