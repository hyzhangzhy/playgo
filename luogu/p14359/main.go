package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

// 先算出值为k的最短区间的l r, 然后再贪心算不相交的l r的最大值

type LR struct {
	L int
	R int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &arr[i])
	}
	mp := map[int]int{}
	sum := 0
	mp[k] = -1
	lr := []LR{}
	for i := 0; i < n; i++ {
		sum = sum ^ arr[i]
		idx, ok := mp[sum]
		if ok {
			lr = append(lr, LR{
				L: idx + 1,
				R: i,
			})
		}
		mp[k^sum] = i
	}
	res := 0
	curR := math.MinInt
	for i := 0; i < len(lr); i++ {
		if lr[i].L <= curR {
			continue
		}
		res++
		curR = lr[i].R
	}
	Fprintf(out, "%d\n", res)
}
