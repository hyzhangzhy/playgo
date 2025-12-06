package main

import (
	"fmt"
	"math/bits"
	"slices"
)

func maxProduct(nums []int) int64 {
	maxm := slices.Max(nums)
	w := bits.Len(uint(maxm))
	u := 1 << w
	uu := u - 1
	dp := make([]int, u)
	for _, x := range nums {
		dp[x] = x
	}

	res := 0
	for s := range dp {
		tmp := s
		for tmp > 0 {
			// 计算最低非0位数
			lowBit := tmp & -tmp
			dp[s] = max(dp[s], dp[s^lowBit])
			tmp ^= lowBit
		}
	}

	for _, x := range nums {
		res = max(res, x*dp[uu^x])
	}

	return int64(res)

}

func main() {
	fmt.Println(maxProduct([]int{1, 2, 3, 4, 5, 6, 7}))

}
