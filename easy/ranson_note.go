package easy

// Given two stings ransomNote and magazine, return true if ransomNote
// can be constructed from magazine and false otherwise.
//
// Each letter in magazine can only be used once in ransomNote.
//
//
//
// Example 1:
//
// Input: ransomNote = "a", magazine = "b"
// Output: false
// Example 2:
//
// Input: ransomNote = "aa", magazine = "ab"
// Output: false
// Example 3:
//
// Input: ransomNote = "aa", magazine = "aab"
// Output: true
func canConstruct(ransomNote string, magazine string) bool {
	chars := make(map[uint8]int)

	for i := 0; i < len(magazine); i++ {
		chars[magazine[i]]++
	}

	for i := 0; i < len(ransomNote); i++ {
		if v, ok := chars[ransomNote[i]]; !ok || v <= 0 {
			return false
		}

		chars[ransomNote[i]]--
	}

	return true
}
