package main

import (
	"fmt"
	"sort"
)

func intersectionSizeTwo(intervals [][]int) int {
	//贪心策略：从小到大排序，从后与2个数，如果不够尽量取最前
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	n := len(intervals)
	res := 2
	cur := []int{intervals[n-1][0], intervals[n-1][0] + 1}
	for i := n - 2; i >= 0; i-- {
		l := intervals[i][0]
		r := intervals[i][1]
		for len(cur) != 0 {
			top := len(cur) - 1
			if r < cur[top] {
				cur = cur[0:top]
			} else {
				break
			}
		}

		if len(cur) == 1 {
			if cur[0] == l {
				cur = append(cur, l+1)
			} else {
				cur = append(cur, l)
				cur[0], cur[1] = cur[1], cur[0]
			}
			res++
		}
		if len(cur) == 0 {
			cur = []int{l, l + 1}
			res += 2
		}
		//fmt.Println(cur)
	}
	return res
}

func main() {
	fmt.Println(intersectionSizeTwo([][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}}))
}
