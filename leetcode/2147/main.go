package main

import "fmt"

func numberOfWays(corridor string) int {
	mod := 1000000007
	tmp := []int{}
	c := 0
	n := len(corridor)
	for i := 0; i < n; i++ {
		if corridor[i] == 'S' {
			c++
		}
	}
	if c%2 == 1 {
		return 0
	}
	if c == 0 {
		return 0
	}
	c = 0
	nxt := 1
	for i := 0; i < n; i++ {
		if corridor[i] == 'S' {
			c++
			if c%2 == 1 {
				tmp = append(tmp, nxt)
				nxt = 1
			}
		}
		if corridor[i] == 'P' {
			if c != 0 && c%2 == 0 {
				nxt++
			}
		}
	}
	res := 1
	for _, v := range tmp {
		res *= v
		res %= mod
	}
	return res
}

func main() {
	fmt.Println(numberOfWays("SSPPSPS"))
}
