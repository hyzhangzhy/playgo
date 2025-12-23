package main

import "sort"

func maxTwoEvents(events [][]int) int {
	n := len(events)
	sort.Slice(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})
	pMax := make([]int, n)
	pMax[0] = events[0][2]
	for i := 1; i < n; i++ {
		pMax[i] = max(pMax[i-1], events[i][2])
	}
	res := events[0][2]
	for i := 1; i < n; i++ {
		tmp := events[i][2]
		left := events[i][0]
		l := 0
		r := i - 1
		for l <= r {
			mid := (l + r) / 2
			if events[mid][1] < left {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		ret := l - 1
		if ret >= 0 {
			tmp += pMax[ret]
		}
		res = max(res, tmp)
	}
	return res
}

func main() {

}
