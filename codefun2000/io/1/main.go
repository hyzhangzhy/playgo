package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	res := 0
	for i := 0; i < n; i++ {
		tmp := 0
		Fscan(in, &tmp)
		res += tmp
	}
	Fprintf(out, "%d", res)

}
