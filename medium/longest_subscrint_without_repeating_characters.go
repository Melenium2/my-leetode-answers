package medium

// Given a string s, find the length of the longest substring without repeating characters.
//
//
//
// Example 1:
//
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
//
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
//
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
//
// Constraints:
//
// 0 <= s.length <= 5 * 104
// s consists of English letters, digits, symbols and spaces.
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		set    = make(map[uint8]struct{})
		maxLen = 1
		index  int
	)

	for len(s) > 0 && index < len(s) {
		curr := s[index]

		if _, ok := set[curr]; ok {
			maxLen = max(maxLen, index)
			s = s[1:]
			index = 0
			prune(set)

			continue
		}

		set[curr] = struct{}{}

		index++
	}

	return max(maxLen, index)
}

func prune(m map[uint8]struct{}) {
	for k := range m {
		delete(m, k)
	}
}

func lengthOfLongestSubstringOptimized(s string) int {
	if len(s) == 0 {
		return 0
	}

	var (
		set    = make(map[uint8]int)
		maxLen = 1
		start  = 0
	)

	for end := 0; end < len(s); end++ {
		curr := s[end]

		if index, ok := set[curr]; ok {
			maxLen = max(maxLen, end-start)
			removeOccurrences(set, s[start:index+1])
			start = index + 1
		}

		set[curr] = end
	}

	return max(maxLen, len(s)-start)
}

func removeOccurrences(m map[uint8]int, keys string) {
	for i := 0; i < len(keys); i++ {
		delete(m, keys[i])
	}
}

func max(a, b int) int {
	if a < b {
		return b
	}

	return a
}
