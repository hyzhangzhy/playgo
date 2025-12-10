package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	var tp int
	Fscan(in, &tp)
	g := make([][]int, n+1)
	if tp == 1 {
		for i := 0; i < n-1; i++ {
			var from, to int
			Fscan(in, &from, &to)
			g[from] = append(g[from], to)
			g[to] = append(g[to], from)
		}
	} else {
		for i := 0; i < n; i++ {
			var to int
			Fscan(in, &to)
			if to != 0 {
				g[to] = append(g[to], i+1)
			}
		}
	}
	for i := 1; i <= n; i++ {
		sort.Ints(g[i])
	}

	var dfs func(cur, from int)
	dfs = func(cur, from int) {
		Fprintf(out, "%d ", cur)
		for _, nxt := range g[cur] {
			if nxt == from {
				continue
			}
			dfs(nxt, cur)
		}
	}
	dfs(1, -1)

}
