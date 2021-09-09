package easy

// Given a string s which consists of lowercase or uppercase letters,
// return the length of the longest palindrome that can be built with those letters.
//
// Letters are case sensitive, for example, "Aa" is not considered a palindrome here.
//
//
//
// Example 1:
//
// Input: s = "abccccdd"
// Output: 7
// Explanation:
// One longest palindrome that can be built is "dccaccd", whose length is 7.
// Example 2:
//
// Input: s = "a"
// Output: 1
// Example 3:
//
// Input: s = "bb"
// Output: 2
func longestPalindrome(s string) int {
	chars := make([]int, 128)

	for i := 0; i < len(s); i++ {
		chars[s[i]]++
	}

	odd := 0

	// count odd chars
	for i := 0; i < len(chars); i++ {
		if chars[i]%2 == 1 {
			odd++
		}
	}

	// if odd chars more than zero than
	// sub all odd chars from len of s
	// and apply 1 because we need 1 uniq char
	// in center of palindrome
	if odd > 0 {
		return len(s) - odd + 1
	} else {
		// if odd chars is zero than it is perfect palindrome,
		// and we just return len of s
		return len(s)
	}
}
