package main

import (
	"bufio"
	. "fmt"
	"os"
)

/*
4 3
2 5 4 3
2 1 3
3 2 4
4 2 4
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	Fscan(in, &n, &m)
	ri := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &ri[i])
	}
	di := make([]int, m+1)
	si := make([]int, m+1)
	ti := make([]int, m+1)
	for i := 1; i <= m; i++ {
		Fscan(in, &di[i], &si[i], &ti[i])
	}

	diff := make([]int, n+2)
	prev := 0
	for i := 1; i <= n; i++ {
		diff[i] = ri[i] - prev
		prev = ri[i]
	}
	check := func(x int) bool {
		// 差分数组
		for i := 0; i <= x; i++ {
			l := si[i]
			r := ti[i]
			sub := di[i]
			diff[l] -= sub
			diff[r+1] += sub
		}
		res := true
		sum := 0
		for i := 1; i <= n; i++ {
			sum += diff[i]
			if sum < 0 {
				res = false
				break
			}
		}
		// 加回来
		for i := 0; i <= x; i++ {
			l := si[i]
			r := ti[i]
			sub := di[i]
			diff[l] += sub
			diff[r+1] -= sub
		}
		return res
	}
	l := 0
	r := m - 1
	for l <= r {
		mid := (l + r) >> 1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	result := r
	if result == m-1 {
		Fprintf(out, "%d\n", 0)
	} else {
		Fprintf(out, "%d\n%d\n", -1, result+1)
	}

}
