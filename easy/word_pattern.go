package easy

import "strings"

// Given a pattern and a string s, find if s follows the same pattern.
//
// Here follow means a full match, such that there is a bijection between a
// letter in pattern and a non-empty word in s.
//
//
//
// Example 1:
// Input: pattern = "abba", s = "dog cat cat dog"
// Output: true
//
// Example 2:
// Input: pattern = "abba", s = "dog cat cat fish"
// Output: false
//
// Example 3:
// Input: pattern = "aaaa", s = "dog cat cat dog"
// Output: false
//
// Example 4:
// Input: pattern = "abba", s = "dog dog dog dog"
// Output: false
func wordPattern(pattern string, s string) bool {
	splitedS := strings.Split(s, " ")

	if len(pattern) != len(splitedS) {
		return false
	}

	memo := make(map[uint8]string)

	for i := 0; i < len(splitedS); i++ {
		if str, ok := memo[pattern[i]]; ok && str != splitedS[i] {
			return false
		}

		for k, v := range memo {
			if splitedS[i] == v && pattern[i] != k {
				return false
			}
		}

		memo[pattern[i]] = splitedS[i]
	}

	return true
}