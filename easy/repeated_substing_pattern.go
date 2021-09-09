package easy

// Given a string s, check if it can be constructed by taking a substring of it
// and appending multiple copies of the substring together.
//
//
//
// Example 1:
//
// Input: s = "abab"
// Output: true
// Explanation: It is the substring "ab" twice.
// Example 2:
//
// Input: s = "aba"
// Output: false
// Example 3:
//
// Input: s = "abcabcabcabc"
// Output: true
// Explanation: It is the substring "abc" four times or the substring "abcabc" twice.
func repeatedSubstringPattern(s string) bool {
	if len(s) < 2 {
		return false
	}

	pattern := ""
	patternLen := 0
	lenS := len(s)

	for i := lenS / 2; i >= 1; i-- {
		if lenS%i == 0 {
			pattern = s[:i]
			patternLen = lenS / i
			str := ""

			for j := 0; j < patternLen; j++ {
				str += pattern
			}

			if str == s {
				return true
			}
		}
	}

	return false
}
