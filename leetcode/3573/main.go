package main

import (
	"fmt"
	"math"
)

func maximumProfit(prices []int, k int) int64 {
	n := len(prices)
	dp := make([][][3]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([][3]int, k+2)
	}
	for i := 0; i <= n; i++ {
		for j := 0; j <= k+1; j++ {
			dp[i][j][1] = math.MinInt / 2
			dp[i][j][2] = math.MinInt / 2
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= k; j++ {
			dp[i+1][j+1][0] = max(dp[i][j+1][0], max(dp[i][j][1]+prices[i], dp[i][j][2]-prices[i]))
			dp[i+1][j+1][1] = max(dp[i][j+1][1], dp[i][j+1][0]-prices[i])
			dp[i+1][j+1][2] = max(dp[i][j+1][2], dp[i][j+1][0]+prices[i])
		}
	}
	return int64(dp[n][k+1][0])
}

func main() {
	fmt.Println(maximumProfit([]int{1, 7, 9, 8, 2}, 2))
}
