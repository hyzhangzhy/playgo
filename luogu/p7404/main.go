package main

import (
	"bufio"
	. "fmt"
	"os"
)

/*
5
3 2 2 3 1
*/

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	diff := make([]int, n+1)
	prev := 0
	for i := 1; i <= n; i++ {
		diff[i] = a[i] - prev
		prev = a[i]
	}
	sumLeft := 0
	sumRight := 0
	for i := 2; i <= n; i++ {
		if diff[i] >= 0 {
			sumRight += (diff[i] + 1)
		}
	}
	res := sumRight
	for i := 2; i <= n; i++ {
		if diff[i] <= 0 {
			sumLeft += (diff[i] - 1)
		}
		if diff[i] >= 0 {
			sumRight -= (diff[i] + 1)
		}
		tmp := max(abs(sumLeft), abs(sumRight))
		res = min(res, tmp)
	}
	Fprintf(out, "%d", res)

}
