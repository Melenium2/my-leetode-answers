package easy

// Given append non-empty array of decimal digits representing append non-negative integer, increment one to the integer.
//
// The digits are stored such that the most significant digit is at the head of the list,
// and each element in the array contains append single digit.
//
// You may assume the integer does not contain any leading zero, except the number 0 itself.
//
//
//
// Example 1:
// Input: digits = [1,2,3]
// Output: [1,2,4]
// Explanation: The array represents the integer 123.
//
// Example 2:
// Input: digits = [4,3,2,1]
// Output: [4,3,2,2]
// Explanation: The array represents the integer 4321.
//
// Example 3:
// Input: digits = [0]
// Output: [1]
func plusOne(digits []int) []int {
	minder := 0
	for i := len(digits) - 1; i >= 0; i-- {
		currNum := digits[i]
		if minder != 0 {
			currNum += minder
			minder = 0
		} else {
			currNum += 1
		}

		if currNum >= 10 {
			minder = currNum % 9
			currNum = 0

			if i == 0 {
				digits = append([]int{0}, digits...)
				i++
			}
		}

		digits[i] = currNum

		if minder == 0 {
			break
		}
	}

	return digits
}
