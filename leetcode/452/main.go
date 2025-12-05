package main

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	n := len(points)
	res := 1
	left := points[n-1][0]
	for i := n - 2; i >= 0; i-- {
		l := points[i][0]
		r := points[i][1]
		if l <= left && left <= r {
			continue
		} else {
			left = l
			res++
		}
	}
	return res
}

func main() {
	fmt.Println(findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}))
}
