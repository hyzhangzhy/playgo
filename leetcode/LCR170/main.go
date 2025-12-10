package main

import (
	"fmt"
	"sort"
)

const fenwickInitVal = 0 // -1e18

type fenwick []int

func newFenwickTree(n int) fenwick {
	t := make(fenwick, n+1)
	for i := range t {
		t[i] = fenwickInitVal
	}
	return t
}

func (fenwick) op(a, b int) int {
	return a + b // max(a, b)
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

// 求区间和 a[l] + ... + a[r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int) int {
	if r < l {
		return 0
	}
	return f.pre(r) - f.pre(l-1)
}

func reversePairs(record []int) int {
	n := len(record)
	cp := make([]int, n)
	copy(cp, record)
	sort.Ints(cp)
	mp := map[int]int{}
	k := 1
	for _, v := range cp {
		_, ok := mp[v]
		if !ok {
			mp[v] = k
			k++
		}
	}
	res := 0
	ftree := newFenwickTree(k)
	for i := 1; i < n; i++ {
		ftree.update(mp[record[i-1]], 1)
		// 加上前面大于当前的数
		res += ftree.query(mp[record[i]]+1, k)
	}
	return res
}

func main() {
	fmt.Println(reversePairs([]int{9, 7, 5, 4, 6}))
}
