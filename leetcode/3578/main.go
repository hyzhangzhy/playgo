package main

import "fmt"

// 单调队列优化动态规划
// 单调队列可以用来类似滑窗的解决区间最大最小问题, O(1)的递推
func countPartitions(nums []int, k int) int {
	mod := 1000000007
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 1
	maxQue := []int{}
	minQue := []int{}
	sum := 0
	left := 1
	for i := 1; i <= n; i++ {
		for len(maxQue) != 0 {
			top := len(maxQue) - 1
			if nums[maxQue[top]] <= nums[i-1] {
				maxQue = maxQue[0:top]
			} else {
				break
			}
		}
		for len(minQue) != 0 {
			top := len(minQue) - 1
			if nums[minQue[top]] >= nums[i-1] {
				minQue = minQue[0:top]
			} else {
				break
			}
		}
		maxQue = append(maxQue, i-1)
		minQue = append(minQue, i-1)
		for {
			max1 := nums[maxQue[0]]
			min1 := nums[minQue[0]]
			if max1-min1 <= k {
				break
			} else {
				sum -= dp[left-1]
				sum = (sum + mod) % mod
				if left-1 == maxQue[0] {
					maxQue = maxQue[1:]
				}
				if left-1 == minQue[0] {
					minQue = minQue[1:]
				}
				left++
			}
		}
		sum += dp[i-1]
		sum %= mod
		dp[i] = sum

	}
	//fmt.Println(dp)
	return dp[n]
}

func main() {
	fmt.Println(countPartitions([]int{9, 4, 1, 3, 7}, 4))
}
