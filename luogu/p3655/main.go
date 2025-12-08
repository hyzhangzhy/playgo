package main

import (
	"bufio"
	. "fmt"
	"os"
)

/*
4 3 2 3
0
5
2
4
6
1 2 1
3 4 -3
1 4 2
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, q, s, t int
	Fscan(in, &n, &q, &s, &t)
	arr := make([]int, n+1)
	for i := 0; i <= n; i++ {
		Fscan(in, &arr[i])
	}
	diff := make([]int, n+2)
	prev := 0
	for i := 0; i <= n; i++ {
		diff[i] = arr[i] - prev
		prev = arr[i]
	}
	sum := 0
	for i := 1; i <= n; i++ {
		if diff[i] < 0 {
			sum += -diff[i] * t
		}
		if diff[i] > 0 {
			sum += -diff[i] * s
		}
	}
	for i := 0; i < q; i++ {
		var x, y, z int
		Fscan(in, &x, &y, &z)
		oldx := diff[x]
		newx := oldx + z

		if oldx <= 0 && newx <= 0 {
			sum += -z * t
		} else if oldx >= 0 && newx >= 0 {
			sum += -z * s
		} else if oldx < 0 && newx > 0 {
			sum += t*oldx - newx*s
		} else if oldx > 0 && newx < 0 {
			sum += s*oldx - newx*t
		}
		diff[x] += z

		if y+1 <= n {
			oldx := diff[y+1]
			z = -z
			newx := oldx + z
			if oldx <= 0 && newx <= 0 {
				sum += -z * t
			} else if oldx >= 0 && newx >= 0 {
				sum += -z * s
			} else if oldx < 0 && newx > 0 {
				sum += t*oldx - newx*s
			} else if oldx > 0 && newx < 0 {
				sum += s*oldx - newx*t
			}
			diff[y+1] += z
		}
		Fprintf(out, "%d\n", sum)
	}
}
