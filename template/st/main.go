package main

import (
	"fmt"
	"math/bits"
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

func main() {
	nums := []int{3, 1, 4, 1, 5, 9, 2, 6}
	stb := newSparseTable(nums, func(a, b int) int { return max(a, b) })
	fmt.Println(stb.query(0, 0))
	fmt.Println(stb.query(0, 5))
	fmt.Println(stb.query(0, 4))
}
