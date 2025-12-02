package main

func checkInclusion(s1 string, s2 string) bool {
	mp := make([]int, 26)
	for i := 0; i < len(s1); i++ {
		mp[s1[i]-'a']++
	}
	res := false
	cntWnd := make([]int, 26)
	l := 0
	for i := 0; i < len(s2); i++ {
		cur := int(s2[i] - 'a')
		if mp[cur] == 0 {
			cntWnd = make([]int, 26)
			l = i + 1
		} else {
			cntWnd[cur]++
			if cntWnd[cur] > mp[cur] {
				for l < i && cntWnd[cur] > mp[cur] {
					cntWnd[s2[l]-'a']--
					l++
				}
			}
			ok := true
			for i := 0; i < 26; i++ {
				if cntWnd[i] != mp[i] {
					ok = false
					break
				}
			}
			if ok {
				res = true
				break
			}
		}

	}
	return res
}

func main() {

}
