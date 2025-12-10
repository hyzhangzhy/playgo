package main

import (
	"fmt"
	"math"
)

func countPermutations(complexity []int) int {
	count := 0
	mod := 1000000007
	minm := math.MaxInt
	n := len(complexity)
	for i := 0; i < n; i++ {
		minm = min(minm, complexity[i])
	}
	if complexity[0] != minm {
		return 0
	}
	for i := 0; i < n; i++ {
		if complexity[i] == minm {
			count++
		}
	}
	if count > 1 {
		return 0
	}
	n = n - 1
	res := 1
	for n > 0 {
		res *= n
		res %= mod
		n--
	}
	return res

}

func main() {
	fmt.Println(countPermutations([]int{3, 3, 3, 4, 4, 4}))
}
