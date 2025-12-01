package main

import (
	"fmt"
	"slices"
	"strings"
)

func reverseWords(s string) string {
	strs := strings.Split(s, " ")
	strs1 := []string{}
	for _, v := range strs {
		if len(v) != 0 {
			strs1 = append(strs1, v)
		}
	}
	slices.Reverse(strs1)
	return strings.Join(strs1, " ")
}

func main() {
	fmt.Println(reverseWords("a good   example"))
}
