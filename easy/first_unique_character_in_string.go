package easy

// Given a string s, find the first non-repeating character in it and return its index. If it does not exist, return -1.
//
//
//
// Example 1:
// Input: s = "leetcode"
// Output: 0
//
// Example 2:
// Input: s = "loveleetcode"
// Output: 2
//
// Example 3:
// Input: s = "aabb"
// Output: -1
func firstUniqChar(s string) int {
	chars := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		chars[s[i]]++
	}

	for i := 0; i < len(s); i++ {
		if chars[s[i]] == 1 {
			return i
		}
	}

	return -1
}
