package main

import "fmt"

func subsets(nums []int) [][]int {
	var result [][]int
	n := len(nums)
	var dfs func(st int)
	tmp := []int{}
	// 0 1 3 5 3
	// 0
	dfs = func(st int) {
		cp := make([]int, len(tmp))
		copy(cp, tmp)
		result = append(result, cp)
		for i := st; i < n; i++ {
			tmp = append(tmp, nums[i])
			dfs(i + 1)
			tmp = tmp[0 : len(tmp)-1]
		}
	}

	dfs(0)
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}
