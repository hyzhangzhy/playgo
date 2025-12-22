package main

import (
	"fmt"
	"math"
	"slices"
)

func minNumberOfSeconds(mountainHeight int, workerTimes []int) int64 {
	maxm := slices.Max(workerTimes)
	n := len(workerTimes)
	hSingle := (mountainHeight-1)/n + 1
	l := 1
	r := maxm * (hSingle) * (hSingle + 1) / 2
	check := func(x int) bool {
		l := mountainHeight
		for _, t := range workerTimes {
			l -= (-1 + int(math.Sqrt(float64(x/t*8+1)))) / 2
			if l <= 0 {
				return true
			}
		}
		return false
	}
	for l <= r {
		mid := l + (r-l)/2
		if check(mid) {
			r = mid - 1
		} else {
			l = mid + 1
		}

	}
	return int64(r + 1)
}

func main() {
	fmt.Println(minNumberOfSeconds(10, []int{3, 2, 2, 4}))
}
