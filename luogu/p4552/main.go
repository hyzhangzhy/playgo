package main

import (
	"bufio"
	. "fmt"
	"os"
)

/*
4
1
1
2
2
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		Fscan(in, &arr[i])
	}
	diff := make([]int, n+1)
	prev := 0
	for i := 0; i < n; i++ {
		diff[i] = arr[i] - prev
		prev = arr[i]
	}
	sum1 := 0
	sum2 := 0
	for i := 1; i < n; i++ {
		if diff[i] > 0 {
			sum1 += diff[i]
		}
		if diff[i] < 0 {
			sum2 += diff[i]
		}
	}
	sum2 = -sum2
	minmOp := max(sum1, sum2)
	count := max(sum1, sum2) - min(sum1, sum2) + 1
	Fprintf(out, "%d\n%d\n", minmOp, count)
}
