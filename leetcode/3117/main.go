package main

import "fmt"

func minimumValueSum(nums []int, andValues []int) int {
	//inf := math.MaxInt / 2
	n := len(nums)
	m := len(andValues)
	dp := make([]map[int]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = map[int]int{}
	}
	dp[0][-1] = 0
	for i := 0; i < n; i++ {
		cur := nums[i]
		newDp := make([]map[int]int, m+1)
		for i := 0; i <= m; i++ {
			newDp[i] = map[int]int{}
		}
		for j := 1; j <= m; j++ {
			for k := range dp[j-1] {
				if k&cur == andValues[j-1] {
					_, ok := newDp[j][-1]
					if !ok {
						newDp[j][-1] = dp[j-1][k] + cur
					} else {
						newDp[j][-1] = min(newDp[j][-1], dp[j-1][k]+cur)
					}
					// _, ok = newDp[j][-1]
					// if !ok {
					// 	newDp[j][-1] = dp[j-1][k] + cur
					// } else {
					// 	newDp[j][-1] = min(newDp[j][-1], dp[j-1][k]+cur)
					// }
				}
				_, ok := newDp[j-1][k&cur]
				if !ok {
					newDp[j-1][k&cur] = dp[j-1][k]
				} else {
					newDp[j-1][k&cur] = min(newDp[j-1][k&cur], dp[j-1][k])
				}

			}
		}
		//fmt.Println(i)
		//fmt.Println(dp)
		dp = newDp
		//dp[0][-1] = 0
		//fmt.Println(dp)
		//fmt.Println("------------------")

	}

	val, ok := dp[m][-1]
	if !ok {
		return -1
	} else {
		return val
	}
}

func main() {
	fmt.Println(minimumValueSum([]int{1, 3, 2, 4, 7, 5, 3}, []int{0, 5, 3}))
}
