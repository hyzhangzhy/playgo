package main

import "fmt"

func specialTriplets(nums []int) int {
	mod := 1000000007
	mpRight := map[int]int{}
	mpLeft := map[int]int{}
	n := len(nums)
	for i := 0; i < n; i++ {
		mpRight[nums[i]]++
	}
	res := 0
	mpRight[nums[0]]--
	for i := 1; i < n-1; i++ {
		mpLeft[nums[i-1]]++
		cur := nums[i] * 2
		mpRight[nums[i]]--
		val1 := mpLeft[cur]
		val2 := mpRight[cur]
		res += val1 * val2
		res %= mod
	}
	return res
}

func main() {
	fmt.Println(specialTriplets([]int{8, 4, 2, 8, 4}))
}
