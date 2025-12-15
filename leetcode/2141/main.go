package main

import "fmt"

func maxRunTime(n int, batteries []int) int64 {
	l := 1
	r := 0
	nb := len(batteries)
	for i := 0; i < nb; i++ {
		r += batteries[i]
	}
	check := func(x int) bool {
		sum := 0
		for i := 0; i < nb; i++ {
			sum += min(batteries[i], x)
		}
		return sum >= n*x
	}
	for l <= r {
		mid := (l + r) / 2
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}

	}
	return int64(l - 1)
}

func main() {
	fmt.Println(maxRunTime(2, []int{3, 3, 3}))
}
