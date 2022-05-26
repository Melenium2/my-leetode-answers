package hard

import "math"

// Given two strings s and t of lengths m and n respectively, return the minimum
// window substring of s such that every character in t (including duplicates) is
// included in the window. If there is no such substring, return the empty string "".
//
// The testcases will be generated such that the answer is unique.
//
// A substring is a contiguous sequence of characters within the string.
//
//
//
// Example 1:
//
// Input: s = "ADOBECODEBANC", t = "ABC"
// Output: "BANC"
// Explanation: The minimum window substring "BANC" includes 'A', 'B', and 'C' from string t.
//
// Example 2:
//
// Input: s = "a", t = "a"
// Output: "a"
// Explanation: The entire string s is the minimum window.
//
// Example 3:
//
// Input: s = "a", t = "aa"
// Output: ""
// Explanation: Both 'a's from t must be included in the window.
// Since the largest window of s only has one 'a', return empty string.
//
//
// Constraints:
//
// m == s.length
// n == t.length
// 1 <= m, n <= 105
// s and t consist of uppercase and lowercase English letters.
//
//
// Follow up: Could you find an algorithm that runs in O(m + n) time?
func minWindow(s string, t string) string {
	if len(t) == 0 {
		return ""
	}

	tSet := make(map[uint8]int, len(t))
	window := make(map[uint8]int)

	for i := 0; i < len(t); i++ {
		curr := t[i]

		tSet[curr]++
	}

	var (
		indexes = make([]int, 2)
		resLen  = math.MaxInt

		have, need = 0, len(tSet)
		start      = 0
	)

	for end := 0; end < len(s); end++ {
		curr := s[end]

		window[curr]++

		necessary, ok := tSet[curr]
		if ok && window[curr] == necessary {
			have++
		}

		for have == need {
			if (end - start + 1) < resLen {
				indexes[0], indexes[1] = start, end
				resLen = end - start + 1
			}

			window[s[start]]--

			necessary, ok = tSet[s[start]]
			if ok && window[s[start]] < necessary {
				have--
			}

			start++
		}
	}

	if resLen == math.MaxInt {
		return ""
	}

	l, r := indexes[0], indexes[1]

	return s[l : r+1]
}
