package main

import "fmt"

/*
分治，思路1: 自上向下
奇数+偶数=奇数，奇数不能被2整除
按下标是奇数还是偶数分组，分组之间都满足条件
向下递归每i步干掉公差是2^(i-1)的排序
*/
func beautifulArray(n int) []int {
	//2 1 4 3
	res := make([]int, n)
	for i := 0; i < n; i++ {
		res[i] = i + 1
	}
	var dfs func(arr []int, l, r int)
	dfs = func(arr []int, l, r int) {
		if l == r {
			return
		}
		tmp := make([]int, r-l+1)
		copy(tmp, arr[l:r+1])
		k := 0
		k1 := 0
		n := len(tmp)
		st := n / 2
		for i := 0; i < len(tmp); i++ {
			if i%2 == 1 {
				arr[l+k] = tmp[i]
				k++
			} else {
				arr[l+k1+st] = tmp[i]
				k1++
			}
		}
		dfs(arr, l, l+st-1)
		dfs(arr, l+st, r)
	}
	dfs(res, 0, n-1)
	return res
}

func main() {
	fmt.Println(beautifulArray(5))
}
