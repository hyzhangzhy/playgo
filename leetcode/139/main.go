package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	mp := map[string]struct{}{}
	for _, v := range wordDict {
		mp[v] = struct{}{}
	}
	n := len(s)
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = -1
	}
	var dfs func(idx int) bool
	dfs = func(idx int) bool {
		if idx == 0 {
			return true
		}
		if dp[idx] != -1 {
			if dp[idx] == 1 {
				return true
			} else {
				return false
			}
		}
		res := false
		for i := idx - 1; i >= 0; i-- {
			sub := s[i:idx]
			_, ok := mp[sub]
			if ok {
				res = res || dfs(i)
			}
		}
		if res {
			dp[idx] = 1
		} else {
			dp[idx] = 0
		}
		return res
	}

	return dfs(n)
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
