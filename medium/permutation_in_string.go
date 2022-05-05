package medium

// Given two strings s1 and s2, return true if s2 contains a permutation of s1, or false otherwise.
//
// In other words, return true if one of s1's permutations is the substring of s2.
//
//
//
// Example 1:
//
// Input: s1 = "ab", s2 = "eidbaooo"
// Output: true
// Explanation: s2 contains one permutation of s1 ("ba").
//
// Example 2:
//
// Input: s1 = "ab", s2 = "eidboaoo"
// Output: false
//
//
// Constraints:
//
// 1 <= s1.length, s2.length <= 104
// s1 and s2 consist of lowercase English letters.
func checkPermutation(s1 string, s2 string) bool {
	if len(s1) == 0 || len(s2) == 0 {
		return false
	}

	var (
		startIndex = 0
		expected   [26]uint8
	)

	for i := 0; i < len(s1); i++ {
		expected[s1[i]-'a']++
	}

	for endIndex := len(s1); endIndex <= len(s2); endIndex++ {
		var runes [26]uint8
		substring := s2[startIndex:endIndex]
		nextIter := false

		for i := 0; i < len(substring); i++ {
			curr := substring[i]
			currIndex := int(curr - 'a')

			if expected[currIndex] == 0 {
				nextIter = true

				break
			}

			runes[currIndex]++

			if runes[currIndex] > expected[currIndex] {
				nextIter = true

				break
			}
		}

		startIndex++

		if !nextIter {
			return true
		}
	}

	return false
}
