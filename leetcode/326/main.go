package main

import "fmt"

func removeDuplicateLetters(s string) string {
	mp := map[byte]int{}
	n := len(s)
	for i := 0; i < n; i++ {
		mp[s[i]]++
	}
	st := []byte{}
	used := make([]bool, 26)
	for i := 0; i < n; i++ {
		use := used[s[i]-'a']
		if !use {
			for len(st) != 0 {
				top := len(st) - 1
				cur := st[top]
				if cur > s[i] {
					if mp[cur] == 0 {
						break
					}

					used[cur-'a'] = false
					st = st[0:top]
				} else {
					break
				}
			}
			st = append(st, s[i])
			used[s[i]-'a'] = true
		}
		mp[s[i]]--

	}
	return string(st)

}

func main() {
	fmt.Println(removeDuplicateLetters("cbacdcbc"))
}
