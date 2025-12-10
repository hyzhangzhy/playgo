package main

import (
	"fmt"
	"sort"
)

func minProductSum(nums1 []int, nums2 []int) int {
	//5 3 4 2
	//4 2 2 5
	//dp[i] = min(dp[i - 1] + 2 * 5)
	//6*6 + 7*7
	//6*7 + 7*6
	//2 3 4 6 8
	//7 5 4 2 1
	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	sort.Slice(nums2, func(i, j int) bool {
		return nums2[i] > nums2[j]
	})
	res := 0
	for i := 0; i < len(nums1); i++ {
		res += nums1[i] * nums2[i]
	}
	return res
}

func main() {
	fmt.Println(minProductSum([]int{2, 1, 4, 5, 7}, []int{3, 2, 4, 8, 6}))
}
