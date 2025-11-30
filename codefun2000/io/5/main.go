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
	mat1 := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		mat1[i] = make([]bool, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var tmp int
			Fscan(in, &tmp)
			if tmp == 1 {
				mat1[i][j] = true
			}
		}
	}
	mat2 := make([][]bool, n+1)
	for i := 0; i < n+1; i++ {
		mat2[i] = make([]bool, n+1)
	}
	for i := 0; i < n; i++ {
		var from, k int
		Fscan(in, &from, &k)
		for j := 0; j < k; j++ {
			var to int
			Fscan(in, &to)
			mat2[from][to] = true
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if mat1[i][j] != mat2[i][j] {
				Fprintln(out, "NO")
				return
			}
		}
	}
	Fprintln(out, "YES")

}
