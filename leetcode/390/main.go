package main

import "fmt"

func lastRemaining(n int) int {
	var dfs func(x int, isRight bool) int
	dfs = func(x int, isRight bool) int {
		if x == 1 {
			return 1
		}
		if !isRight {
			return 2 * dfs(x/2, !isRight)
		} else {
			if x%2 != 0 {
				return 2 * dfs(x/2, !isRight)
			} else {
				return 2*dfs(x/2, !isRight) - 1
			}
		}
	}
	return dfs(n, false)
}

func main() {
	fmt.Println(lastRemaining(25))
}
