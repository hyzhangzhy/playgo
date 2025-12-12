package main

import (
	"fmt"
	"sort"
)

const fenwickInitVal = -1 // -1e18

type fenwick []int

func newFenwickTree(n int) fenwick {
	t := make(fenwick, n+1)
	for i := range t {
		t[i] = fenwickInitVal
	}
	return t
}

func (fenwick) op(a, b int) int {
	return max(a, b)
}

// a[i] 增加 val
// 1<=i<=n
func (f fenwick) update(i, val int) {
	for ; i < len(f); i += i & -i {
		f[i] = f.op(f[i], val)
	}
}

// 求前缀和 a[1] + ... + a[i]
// 1<=i<=n
func (f fenwick) pre(i int) int {
	res := fenwickInitVal
	i = min(i, len(f)-1)
	for ; i > 0; i &= i - 1 {
		res = f.op(res, f[i])
	}
	return res
}

type P struct {
	x     int
	y     int
	count int
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	n := len(nums1)
	m := len(queries)
	p := make([]P, n+m)
	for i := 0; i < n; i++ {
		p[i].x = nums1[i]
		p[i].y = nums2[i]
		// 与下面区分开
		p[i].count = -(nums1[i] + nums2[i])
	}
	for i := 0; i < m; i++ {
		idx := i + n
		p[idx].x = queries[i][0]
		p[idx].y = queries[i][1]
		p[idx].count = i
	}
	sort.Slice(p, func(i, j int) bool {
		if p[i].x == p[j].x {
			if p[i].y == p[j].y {
				return p[i].count < p[j].count
			}
			return p[i].y > p[j].y
		}
		return p[i].x > p[j].x
	})

	yslic := make([]int, n+m)
	for i := 0; i < n+m; i++ {
		yslic[i] = p[i].y
	}
	sort.Slice(yslic, func(i, j int) bool {
		return yslic[i] < yslic[j]
	})
	k := 1
	mp := map[int]int{}
	for i := 0; i < n+m; i++ {
		_, ok := mp[yslic[i]]
		if !ok {
			mp[yslic[i]] = k
			k++
		}
	}
	bt := newFenwickTree(k)
	res := make([]int, m)
	for i := 0; i < n+m; i++ {
		//x := p[i].x
		y := k - mp[p[i].y]
		v := p[i].count
		if v >= 0 {
			res[v] = bt.pre(y)
		} else {
			bt.update(y, -v)
		}
	}
	return res

}

func main() {
	fmt.Println(maximumSumQueries([]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 5}}))
}
