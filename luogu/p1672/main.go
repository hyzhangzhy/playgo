package main

import (
	"bufio"
	. "fmt"
	"os"
)

/*
3 14 4 10
1 9
5 8
8 12
*/
type LR struct {
	L int
	R int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var c, f1, f2, d int
	Fscan(in, &c, &f1, &f2, &d)
	lr := make([]LR, c)
	maxm := 0
	for i := 0; i < c; i++ {
		Fscan(in, &lr[i].L, &lr[i].R)
		maxm = max(maxm, lr[i].R)
	}
	diff := make([]int, maxm+2)
	for i := 0; i < c; i++ {
		diff[lr[i].L] += 1
		diff[lr[i].R+1] -= 1
	}
	count := make([]int, maxm+2)
	prev := 0
	for i := 0; i < maxm+2; i++ {
		count[i] = prev + diff[i]
		prev = count[i]
	}
	if d >= maxm {
		d = maxm
	}
	for i := d; i >= 0; i-- {
		f2 += count[i]
		if f1 == f2 {
			Fprintf(out, "%d\n", i)
			break
		}
	}
	//return -1

}
