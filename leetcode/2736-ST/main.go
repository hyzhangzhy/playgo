package main

import (
	"fmt"
	"math/bits"
	"sort"
)

type sparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

// 模板和readme中相反, i是向后的2^i, j是开始的位置
// 时间复杂度 O(n * log n)
func newSparseTable[T any](a []T, op func(T, T) T) sparseTable[T] {
	n := len(a)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], a)
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable[T]{st, op}
}

// [l,r] 下标从 0 开始
// 返回 op(nums[l:r])
// 时间复杂度 O(1)
func (s sparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l+1)) - 1
	return s.op(s.st[k][l], s.st[k][r-1<<k+1])
}

type P struct {
	x int
	y int
}

func maximumSumQueries(nums1 []int, nums2 []int, queries [][]int) []int {
	p := []P{}
	n := len(nums1)
	for i := 0; i < n; i++ {
		p = append(p, P{
			x: nums1[i],
			y: nums2[i],
		})
	}
	sort.Slice(p, func(i, j int) bool {
		return p[i].x > p[j].x
	})
	p1 := []P{}
	p1 = append(p1, p[0])
	for i := 1; i < n; i++ {
		top := len(p1) - 1
		if p[i].y <= p1[top].y {
			continue
		}
		p1 = append(p1, p[i])
	}
	nq := len(queries)
	np1 := len(p1)
	res := make([]int, nq)
	nums := make([]int, np1)
	for i := 0; i < np1; i++ {
		nums[i] = p1[i].x + p1[i].y
	}
	stb := newSparseTable(nums, func(a, b int) int { return max(a, b) })
	for i := 0; i < nq; i++ {
		x := queries[i][0]
		y := queries[i][1]
		r := sort.Search(np1, func(i int) bool {
			return p1[i].x < x
		}) - 1
		l := sort.Search(np1, func(i int) bool {
			return p1[i].y >= y
		})
		if l > r {
			res[i] = -1
		} else {
			res[i] = stb.query(l, r)
		}
	}
	return res
}

func main() {
	fmt.Println(maximumSumQueries([]int{4, 3, 1, 2}, []int{2, 4, 9, 5}, [][]int{{4, 1}, {1, 3}, {2, 5}}))
}
