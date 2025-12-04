package main

import (
	"fmt"
	"math"
	"sort"
)

// 非O(n)
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	cp := make([]int, n)
	copy(cp, nums)
	sort.Ints(cp)
	l := math.MaxInt
	r := math.MinInt
	for i := 0; i < n; i++ {
		if cp[i] != nums[i] {
			l = min(l, i)
			r = max(r, i)
		}
	}
	if l == math.MaxInt {
		return 0
	}
	return r - l + 1
}

func main() {
	fmt.Println([]int{2, 6, 4, 8, 10, 9, 15})
}
