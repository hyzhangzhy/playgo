package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
乘积最大子数组
*/
func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := []int{}
	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(in, &x)
		arr = append(arr, x)
	}
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0] = arr[0]
	dpMin[0] = arr[0]
	res := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] < 0 {
			dpMax[i] = max(arr[i], arr[i]*dpMin[i-1])
			dpMin[i] = min(arr[i], arr[i]*dpMax[i-1])
		} else {
			dpMax[i] = max(arr[i], arr[i]*dpMax[i-1])
			dpMin[i] = min(arr[i], arr[i]*dpMin[i-1])
		}
		res = max(res, dpMax[i])
	}

	fmt.Fprintln(out, res)
}
