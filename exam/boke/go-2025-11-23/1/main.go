package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
括号生成
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)

	res := []string{}
	tmp := make([]byte, n*2)
	var dfs func(cur, open, close int)
	dfs = func(cur, open, close int) {
		if cur >= 2*n {
			res = append(res, string(tmp))
			return
		}
		if open < n {
			tmp[cur] = '('
			dfs(cur+1, open+1, close)
		}
		if open > close {
			tmp[cur] = ')'
			dfs(cur+1, open, close+1)
		}
	}
	dfs(0, 0, 0)
	fmt.Fprintln(out, res)
}
