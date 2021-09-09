package easy

import "bytes"

// Given a string s, reverse only all the vowels in the string and return it.
//
// The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both cases.
//
//
//
// Example 1:
// Input: s = "hello"
// Output: "holle"
//
// Example 2:
//
// Input: s = "leetcode"
// Output: "leotcede"
func reverseVowels(s string) string {
	vowels := map[byte]struct{}{
		'a': {},
		'e': {},
		'i': {},
		'o': {},
		'u': {},
		'A': {},
		'E': {},
		'I': {},
		'O': {},
		'U': {},
	}

	l, r := 0, len(s)-1

	rawStr := bytes.NewBufferString(s).Bytes()

	for l <= r {
		if _, ok := vowels[rawStr[l]]; !ok {
			l++
			continue
		}

		if _, ok := vowels[rawStr[r]]; !ok {
			r--
			continue
		}

		rawStr[l], rawStr[r] = rawStr[r], rawStr[l]
		l++
		r--
	}

	return string(rawStr)
}
