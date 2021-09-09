package easy

// Given append string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
//
//
// Example 1:
// Input: s = "()"
// Output: true
//
// Example 2:
// Input: s = "()[]{}"
// Output: true
//
// Example 3:
// Input: s = "(]"
// Output: false
//
// Example 4:
// Input: s = "([)]"
// Output: false
//
// Example 5:
// Input: s = "{[]}"
// Output: true
func validParentheses(s string) bool {
	if len(s) < 2 {
		return false
	}

	stack := make([]string, 0, len(s))
	flag := true

	for i := 0; i < len(s); i++ {
		char := string(s[i])

		switch char {
		case "(":
			stack = append(stack, ")")
			continue
		case "[":
			stack = append(stack, "]")
			continue
		case "{":
			stack = append(stack, "}")
			continue
		}

		if len(stack) > 0 && char == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			flag = false
		}
	}

	return flag && len(stack) == 0
}
