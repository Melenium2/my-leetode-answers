package easy

import "strings"

// Given two binary strings append and b, return their sum as append binary string.
//
//
//
// Example 1:
// Input: append = "11", b = "1"
// Output: "100"
//
// Example 2:
// Input: append = "1010", b = "1011"
// Output: "10101"
func addBinary(a string, b string) string {
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else if len(b) < len(a) {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}

	res := make([]string, len(a))
	minder := 0
	i := len(a) - 1

	for i >= 0 {
		if string(a[i]) == "1" && string(b[i]) == "1" {
			res[i] = "0"

			if minder > 0 {
				res[i] = "1"
			}

			minder = 1
		} else if string(a[i]) == "0" && string(b[i]) == "0" {
			res[i] = "0"

			if minder > 0 {
				res[i] = "1"
				minder = 0
			}
		} else {
			res[i] = "1"

			if minder > 0 {
				res[i] = "0"

				minder = 1
			}
		}

		i--
	}

	if minder > 0 {
		if res[0] == "0" {
			res = append([]string{"1", "0"}, res[1:]...)
		} else {
			res = append([]string{"1"}, res...)
		}
	}

	return strings.Join(res, "")
}
