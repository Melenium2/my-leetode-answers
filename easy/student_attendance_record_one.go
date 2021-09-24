package easy

// You are given a string s representing an attendance record for a student where
// each character signifies whether the student was absent, late, or present on that day.
// The record only contains the following three characters:
//
// 'A': Absent.
// 'L': Late.
// 'P': Present.
// The student is eligible for an attendance award if they meet both of the following criteria:
//
// The student was absent ('A') for strictly fewer than 2 days total.
// The student was never late ('L') for 3 or more consecutive days.
// Return true if the student is eligible for an attendance award, or false otherwise.
//
//
//
// Example 1:
// Input: s = "PPALLP"
// Output: true
// Explanation: The student has fewer than 2 absences and was never late 3 or more consecutive days.
//
// Example 2:
// Input: s = "PPALLL"
// Output: false
// Explanation: The student was late 3 consecutive days in the last 3 days, so is not eligible for the award.
func checkRecord(s string) bool {
	absent := 0
	late := 0
	lateFlag := false

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case 'A':
			absent++
			late = 0
		case 'L':
			late++
			if late >= 3 {
				lateFlag = true
			}
		default:
			late = 0
		}
	}

	award := absent >= 2 || lateFlag

	return !award
}
