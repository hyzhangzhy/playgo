package main

import (
	"bufio"
	. "fmt"
	"os"
)

/*
9 10
3 1 4 1 5 9 2 6 5 3
200 100 50
300 299 100
500 200 500
345 234 123
100 50 100
600 100 1
450 400 80
2 1 10
*/

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	pi := make([]int, m)
	for i := 0; i < m; i++ {
		Fscan(in, &pi[i])
	}
	ai := make([]int, n+1)
	bi := make([]int, n+1)
	ci := make([]int, n+1)
	for i := 1; i <= n-1; i++ {
		Fscan(in, &ai[i], &bi[i], &ci[i])
	}

	// 差分数组
	d := make([]int, n+1)

	for i := 0; i < m-1; i++ {
		from := pi[i]
		to := pi[i+1]
		// 1 -> 1-2
		// 2 -> 2-3
		if from > to {
			tmp := from
			from = to
			to = tmp
		}
		d[from] += 1
		d[to] -= 1
	}

	res := 0
	add := 0
	for i := 1; i <= n; i++ {
		add += d[i]
		x1 := ai[i] * add
		x2 := ci[i] + bi[i]*add
		if x1 < x2 {
			res += x1
		} else {
			res += x2
		}
	}
	Fprintf(out, "%d", res)

}
