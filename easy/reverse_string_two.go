package easy

import "bytes"

// Given a string s and an integer k, reverse the first k
// characters for every 2k characters counting from the start of the string.
//
// If there are fewer than k characters left, reverse all of them.
// If there are less than 2k but greater than or equal to k characters, then reverse the first k characters and left the other as original.
//
//
//
// Example 1:
//
// Input: s = "abcdefg", k = 2
// Output: "bacdfeg"
// Example 2:
//
// Input: s = "abcd", k = 2
// Output: "bacd"
func reverseStr(s string, k int) string {
	runes := bytes.NewBufferString(s).Bytes()

	for currIndex := 0; currIndex < len(s); currIndex += 2 * k {
		i := currIndex
		j := min(currIndex+k-1, len(s)-1)
		for i < j {
			runes[i], runes[j] = runes[j], runes[i]
			i++
			j--
		}
	}

	return string(runes)
}
