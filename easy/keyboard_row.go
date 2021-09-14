package easy

// Given an array of strings words, return the words that can be typed using
// letters of the alphabet on only one row of American keyboard like the image below.
//
// In the American keyboard:
//
// the first row consists of the characters "qwertyuiop",
// the second row consists of the characters "asdfghjkl", and
// the third row consists of the characters "zxcvbnm".
//
//
//
// Example 1:
//
// Input: words = ["Hello","Alaska","Dad","Peace"]
// Output: ["Alaska","Dad"]
// Example 2:
//
// Input: words = ["omk"]
// Output: []
// Example 3:
//
// Input: words = ["adsdf","sfd"]
// Output: ["adsdf","sfd"]
func findWords(words []string) []string {
	const upperChars = 32
	alphabet := map[uint8]uint8{
		'q': 1,
		'w': 1,
		'e': 1,
		'r': 1,
		't': 1,
		'y': 1,
		'u': 1,
		'i': 1,
		'o': 1,
		'p': 1,
		'a': 2,
		's': 2,
		'd': 2,
		'f': 2,
		'g': 2,
		'h': 2,
		'j': 2,
		'k': 2,
		'l': 2,
		'z': 3,
		'x': 3,
		'c': 3,
		'v': 3,
		'b': 3,
		'n': 3,
		'm': 3,
	}

	answ := make([]string, 0)

	for _, w := range words {
		chosenRow := uint8(0)

		for i := 0; i < len(w); i++ {
			char := w[i]

			if v, ok := alphabet[char]; ok {
				if chosenRow > 0 && chosenRow != v {
					chosenRow = 0
					break
				}

				chosenRow = v

				continue
			}

			if v, ok := alphabet[char+upperChars]; ok {
				if chosenRow > 0 && chosenRow != v {
					chosenRow = 0
					break
				}

				chosenRow = v
			}
		}

		if chosenRow != 0 {
			answ = append(answ, w)
		}
	}

	return answ
}

// The only thing we need to realise here is the characters of word should be subset of one of the rows,
// and if we map the characters of the word to an integer, then using the same logic - the bits that are
// on in this integer map of word must be on in the integer map of one of the rows to be called a subset
// of it and if it is, it should be added to the answer list.
// now to check if the word is subset of the row i used this logic-
//
// one == (one|cur)
// where one represent integer map of characters in first row, one|cur gives us an integer with bits that are
// on in both row and the word, but if word is subset of row, then the bits that are on in word should already
// be on in the row , hence if this number is equal to the row then word was subset of row.

// public String[] findWords(String[] words) {
//        int one = 0, two = 0, three = 0;
//        for(char c : "qwertyuiop".toCharArray()) one |= (1<<(c-'a'));
//        for(char c : "asdfghjkl".toCharArray()) two |= (1<<(c-'a'));
//        for(char c : "zxcvbnm".toCharArray()) three |= (1<<(c-'a'));
//        List<String> ans = new ArrayList<>();
//        for(String word : words){
//            int cur = 0;
//            for(char c : word.toCharArray()) cur |= (1<<(c-'a'));
//            if(one == (one|cur) || two == (two|cur) || three == (three|cur)) ans.add(word);
//        }
//        return ans.toArray(new String[ans.size()]);
//    }
