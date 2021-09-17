package easy

// We define the usage of capitals in a word to be right when one of the following cases holds:
//
// All letters in this word are capitals, like "USA".
// All letters in this word are not capitals, like "leetcode".
// Only the first letter in this word is capital, like "Google".
// Given a string word, return true if the usage of capitals in it is right.
//
//
//
// Example 1:
// Input: word = "USA"
// Output: true
//
// Example 2:
// Input: word = "FlaG"
// Output: false
func detectCapitalUse(word string) bool {
	capitals := 0
	lower := 0

	for i := 0; i < len(word); i++ {
		if word[i] <= 90 {
			capitals++
		} else {
			lower++
		}

	}

	if lower == len(word) || capitals == len(word) {
		return true
	} else if capitals == 1 && word[0] <= 90 {
		return true
	}

	return false
}
