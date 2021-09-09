package easy

// Given two strings s and t, determine if they are isomorphic.
//
// Two strings s and t are isomorphic if the characters in s can be replaced to get t.
//
// All occurrences of a character must be replaced with another character while
// preserving the order of characters. No two characters may map to the same character,
// but a character may map to itself.
//
//
//
// Example 1:
//
// Input: s = "egg", t = "add"
// Output: true
// Example 2:
//
// Input: s = "foo", t = "bar"
// Output: false
// Example 3:
//
// Input: s = "paper", t = "title"
// Output: true
func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sChars := make([]uint8, 128)
	tChars := make([]uint8, 128)

	for i := 0; i < len(s); i++ {
		if sChars[s[i]] == 0 && tChars[t[i]] == 0 {
			sChars[s[i]] = t[i]
			tChars[t[i]] = s[i]
		} else if sChars[s[i]] != t[i] {
			return false
		}
	}

	return true
}
