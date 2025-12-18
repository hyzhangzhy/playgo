package main

import "fmt"

func maxProfit(prices []int, strategy []int, k int) int64 {
	//1 1 -1 -1
	//1 0 1 1
	//0 0 0 1 1 1
	//  0 0 0 1 1 1
	n := len(prices)
	pSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pSum[i] = pSum[i-1] + prices[i-1]*strategy[i-1]
	}
	res := pSum[n]
	add := 0
	for i := k / 2; i < k; i++ {
		add += prices[i]
	}
	for i := 0; i <= n-k; i++ {
		sub := pSum[i+k] - pSum[i]
		tmp := pSum[n] - sub + add
		res = max(res, tmp)
		if i+k == n {
			break
		}
		add -= prices[i+k/2]
		add += prices[i+k]
	}
	return int64(res)
}

func main() {
	fmt.Println(maxProfit([]int{5, 4, 3}, []int{1, 1, 0}, 2))
}
