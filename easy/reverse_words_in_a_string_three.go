package easy

import "bytes"

// Given a string s, reverse the order of characters in each word within a sentence
// while still preserving whitespace and initial word order.
//
//
//
// Example 1:
// Input: s = "Let's take LeetCode contest"
// Output: "s'teL ekat edoCteeL tsetnoc"

// Example 2:
// Input: s = "God Ding"
// Output: "doG gniD"
//
//
// Constraints:
// 1 <= s.length <= 5 * 104
// s contains printable ASCII characters.
// s does not contain any leading or trailing spaces.
// There is at least one word in s.
// All the words in s are separated by a single space.
func reverseWords(s string) string {
	raw := bytes.NewBufferString(s).Bytes()
	lastWordIndex := 0

	for i := 0; i < len(raw)+1; i++ {
		if i == len(raw) || raw[i] == ' ' {
			start := lastWordIndex
			end := i - 1

			for start < end {
				raw[start], raw[end] = raw[end], raw[start]
				start++
				end--
			}

			lastWordIndex = i + 1
		}
	}

	return string(raw)
}
