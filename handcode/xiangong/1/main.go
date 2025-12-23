package main

import "fmt"

func myAtoi(s string) int32 {
	n := len(s)
	isMinus := false
	res := int32(0)
	i := 0
	for i < n {
		cur := s[i]
		if cur == ' ' {
			i++
			continue
		} else {
			break
		}
	}
	if i < n {
		if s[i] == '-' {
			isMinus = true
			i++

		} else if s[i] == '+' {
			isMinus = false
			i++
		}
	}
	for i < n {
		if s[i] >= '0' && s[i] <= '9' {
			cur := int32(s[i] - '0')
			if res*10 < 0 || res*10+cur < 0 {
				break
			}
			res = res*10 + cur
			i++
		} else {
			break
		}
	}
	if isMinus {
		return -res
	} else {
		return res
	}
}

func main() {
	fmt.Println(myAtoi("    -42"))
	fmt.Println(myAtoi("123"))
	fmt.Println(myAtoi("   -1asdf23 asdfsadfasdf"))
	fmt.Println(myAtoi("22222222222222222222222222222222222222222222666666666666666666666222222222222222222222222"))
}
