package medium

// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
//
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, typically using all the original letters exactly once.
//
//
//
// Example 1:
//
// Input: strs = ["eat","tea","tan","ate","nat","bat"]
// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]
//
// Example 2:
//
// Input: strs = [""]
// Output: [[""]]
//
// Example 3:
//
// Input: strs = ["a"]
// Output: [["a"]]
//
//
// Constraints:
//
// 1 <= strs.length <= 104
// 0 <= strs[i].length <= 100
// strs[i] consists of lowercase English letters.
func groupAnagrams(strs []string) [][]string {
	counts := make([][26]uint8, len(strs))
	groups := make(map[[26]uint8][]string)

	for i := 0; i < len(strs); i++ {
		curr := strs[i]

		for index := 0; index < len(curr); index++ {
			c := curr[index] - 'a'

			counts[i][c]++
		}
	}

	for i := 0; i < len(counts); i++ {
		curr := counts[i]

		groups[curr] = append(groups[curr], strs[i])
	}

	result := make([][]string, 0, len(groups))

	for _, v := range groups {
		result = append(result, v)
	}

	return result
}
