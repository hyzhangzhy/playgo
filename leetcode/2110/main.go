package main

import "fmt"

func getDescentPeriods(prices []int) int64 {
	n := len(prices)
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		if prices[i] == prices[i-1]-1 {
			dp[i] = dp[i-1] + 1
		} else {
			dp[i] = 1
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		res += dp[i]
	}
	return int64(res)
}

func main() {
	fmt.Println(getDescentPeriods([]int{3, 2, 1, 4}))
}
