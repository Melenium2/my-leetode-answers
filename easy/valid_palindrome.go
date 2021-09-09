package easy

import (
	"regexp"
	"strings"
)

// Given a string s, determine if it is a palindrome, considering only
// alphanumeric characters and ignoring cases.
//
//
//
// Example 1:
// Input: s = "A man, a plan, a canal: Panama"
// Output: true
// Explanation: "amanaplanacanalpanama" is a palindrome.
//
// Example 2:
//
// Input: s = "race a car"
// Output: false
// Explanation: "raceacar" is not a palindrome.
func isStringPalindrome(s string) bool {
	palindrome := true

	re := regexp.MustCompile(`[^A-Za-z0-9]`)

	s = strings.ToLower(s)
	s = re.ReplaceAllLiteralString(s, "")

	l, r := 0, len(s)-1

	for l < r {
		lChar := string(s[l])
		rChar := string(s[r])

		if lChar == rChar {
			l++
			r--
		} else {
			palindrome = false

			break
		}
	}

	return palindrome
}
