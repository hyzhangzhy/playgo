package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Solution struct {
	pSum  []int
	total int
}

func Constructor(w []int) Solution {
	n := len(w)
	ps := make([]int, n+1)
	for i := 1; i <= n; i++ {
		ps[i] = ps[i-1] + w[i-1]
	}
	return Solution{
		pSum:  ps,
		total: ps[n],
	}
}

func (this *Solution) PickIndex() int {
	x := rand.Int() % this.total
	check := x + 1
	n := len(this.pSum)
	idx := sort.Search(n, func(x int) bool {
		return this.pSum[x] >= check
	})
	return idx - 1
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */

func main() {
	w := []int{1, 3}
	obj := Constructor(w)
	for i := 0; i < 100; i++ {
		fmt.Println(obj.PickIndex())
	}
}
