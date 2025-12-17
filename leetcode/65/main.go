package main

import (
	"fmt"
	"unicode"
)

func isNumber(s string) bool {
	hasNum := false
	hasDot := false
	hasE := false
	for i, c := range s {
		switch {
		case unicode.IsDigit(c):
			hasNum = true
		case c == '.':
			if hasDot || hasE {
				return false
			}
			hasDot = true
		case c == '+' || c == '-':
			if i != 0 && s[i-1] != 'e' && s[i-1] != 'E' {
				return false
			}
		case c == 'e' || c == 'E':
			if hasE || i == 0 || !hasNum {
				return false
			}
			hasNum = false
			hasE = true

		default:
			return false
		}
	}
	return hasNum

}

func main() {
	fmt.Println(isNumber("."))
}
