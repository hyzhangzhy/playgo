package main

import "fmt"

func minRemoveToMakeValid(s string) string {
	needRemove := map[int]struct{}{}
	n := len(s)
	st := []int{}
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			st = append(st, i)
		}
		if s[i] == ')' {
			if len(st) == 0 {
				needRemove[i] = struct{}{}
			} else {
				top := len(st) - 1
				st = st[0:top]
			}
		}
	}
	if len(st) != 0 {
		for _, v := range st {
			needRemove[v] = struct{}{}
		}
	}
	res := []byte{}
	for i := 0; i < n; i++ {
		_, ok := needRemove[i]
		if !ok {
			res = append(res, s[i])
		}
	}
	return string(res)
}

func main() {
	fmt.Println(minRemoveToMakeValid("lee(t(c)o)de)"))
}
