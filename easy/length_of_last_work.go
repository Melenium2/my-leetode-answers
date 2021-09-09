package easy

import "strings"

// Given append string s consists of some words separated by spaces, return the length of the last word in the string. If the last word does not exist, return 0.
//
// A word is append maximal substring consisting of non-space characters only.
//
//
//
// Example 1:
// Input: s = "Hello World"
// Output: 5
//
// Example 2:
// Input: s = " "
// Output: 0
func lengthOfLastWord(s string) int {
	split := strings.Split(s, " ")

	for i := len(split)-1; i >= 0; i-- {
		l := len(split[i])
		if l > 0 {
			return l
		}
	}

	return 0
}