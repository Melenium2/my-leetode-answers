package easy

// Given two strings s and t, return true if s is a subsequence of t, or false otherwise.
//
// A subsequence of a string is a new string that is formed from the original string by deleting some
// (can be none) of the characters without disturbing the relative positions of the remaining characters.
// (i.e., "ace" is a subsequence of "abcde" while "aec" is not).
//
//
//
// Example 1:
// Input: s = "abc", t = "ahbgdc"
// Output: true
//
// Example 2:
// Input: s = "axc", t = "ahbgdc"
// Output: false
//
//
// Follow up: Suppose there are lots of incoming s, say s1, s2, ..., sk where k >= 10^9,
// and you want to check one by one to see if t has its subsequence. In this scenario,
// how would you change your code?
func isSubsequence(s string, t string) bool {
	index := 0

	for i := 0; i < len(t); i++ {
		if index == len(s) {
			return true
		}

		if t[i] == s[index] {
			index++
		}
	}

	return index == len(s)
}
