package main

import "fmt"

func threeSumMulti(arr []int, target int) int {
	mod := 1000000007
	n := len(arr)
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, 4)
		for j := 0; j <= 3; j++ {
			dp[i][j] = make([]int, target+1)
		}
	}
	dp[0][1][arr[0]] = 1
	for i := 0; i < n; i++ {
		dp[i][0][0] = 1
	}
	for i := 1; i < n; i++ {
		for j := 1; j <= 3; j++ {
			for k := 0; k <= target; k++ {
				dp[i][j][k] = dp[i-1][j][k]
				if k-arr[i] >= 0 {
					dp[i][j][k] += dp[i-1][j-1][k-arr[i]]
					dp[i][j][k] %= mod
				}
			}

		}
	}
	return dp[n-1][3][target]
}

func main() {
	fmt.Println(threeSumMulti([]int{1, 1, 2, 2, 3, 3, 4, 4, 5, 5}, 8))
}
