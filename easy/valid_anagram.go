package easy

// Given two strings s and t, return true if t is an anagram of s, and false otherwise.
//
//
//
// Example 1:
// Input: s = "anagram", t = "nagaram"
// Output: true
//
// Example 2:
// Input: s = "rat", t = "car"
// Output: false
//
// Constraints:
//
// 1 <= s.length, t.length <= 5 * 104
// s and t consist of lowercase English letters.
//
//
// Follow up: What if the inputs contain Unicode characters? How would you adapt your solution to such a case?
// Other solution
//
// We create slice on 26 elements
// Then we iterate over two strings and at char[i] - 'a' increment or decrement value at char of first string
// or second
// Then we iterate over slice of 26 elements and check that all values equals zero
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}


	sChars := make([]int, 128)
	tChars := make([]int, 128)

	for i := 0; i < len(s); i++ {
		sChars[s[i]]++
		tChars[t[i]]++
	}

	for i := 0; i < len(sChars); i++ {
		if sChars[i] != tChars[i] {
			return false
		}
	}

	return true
}