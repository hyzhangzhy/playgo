package main

import (
	"fmt"
	"math"
)

func findUnsortedSubarray(nums []int) int {
	// O(n) 分成3组, sa, sb, sc 其中sa, sc 有序，需要排序sb
	// nums_s[i] <= min[j=(i+1, n - 1)] nums_s[j]
	// 第一个不成立的i
	// nums_s[i] >= max....
	n := len(nums)
	left := n
	right := -1
	maxm := math.MinInt
	minm := math.MaxInt
	for i := 0; i < n; i++ {
		if n-1-i >= 0 && nums[n-1-i] > minm {
			left = n - 1 - i
		}
		if i < n && nums[i] < maxm {
			right = i
		}
		minm = min(minm, nums[n-1-i])
		maxm = max(maxm, nums[i])
	}

	if left == n {
		return 0
	}
	return right - left + 1

}

func main() {
	fmt.Println(findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15}))
}
