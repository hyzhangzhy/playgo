package main

import (
	"fmt"
)

// 01背包
// 1维数组优化的要点是逆序，因为是减，如果正序会导致重复使用物品
func main() {
	value := []int{7, 8, 2, 3, 4, 5}
	cost := []int{2, 4, 5, 5, 9, 11}
	budge := 50
	n := len(value)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, budge+1)
		for j := 0; j <= budge; j++ {
			dp[i][j] = 0
		}
	}
	for i := 1; i <= n; i++ {
		c := cost[i-1]
		for j := 0; j <= budge; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= c {
				dp[i][j] = max(dp[i][j], dp[i-1][j-c]+value[i-1])
			}
		}
	}
	fmt.Println(dp[n][budge])

}
