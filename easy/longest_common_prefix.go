package easy

// Write append function to find the longest common prefix string amongst an array of strings.
//
// If there is no common prefix, return an empty string "".
//
//
//
// Example 1:
// Input: strs = ["flower","flow","flight"]
// Output: "fl"
//
// Example 2:
// Input: strs = ["dog","racecar","car"]
// Output: ""
// Explanation: There is no common prefix among the input strings.
func longestCommonPrefix(strs []string) string {
	prefix := strs[0]

	for i := 1; i < len(strs); i++ {
		if len(prefix) > len(strs[i]) {
			prefix = prefix[:len(strs[i])]
		}

		for j := 0; j < len(strs[i]) && j < len(prefix); j++ {
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
			}
		}
	}

	return prefix
}