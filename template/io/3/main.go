package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(nil, math.MaxInt)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	in.Scan()
	line := in.Bytes()
	sp := bytes.Split(line, []byte{' '})
	// ...
	nums := make([]int, 0, len(sp))
	for _, s := range sp {
		if len(s) == 0 {
			continue
		}
		if num, err := strconv.Atoi(string(s)); err == nil {
			nums = append(nums, num)
		}
	}

	type Id struct {
		id  int
		val int
	}
	g := make([][]Id, len(nums))
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] == -1 {
			continue
		}
		l := i*2 + 1
		r := i*2 + 2
		if l < n && nums[l] != -1 {
			g[i] = append(g[i], Id{
				id:  l,
				val: nums[l],
			})
		}
		if r < n && nums[r] != -1 {
			g[i] = append(g[i], Id{
				id:  r,
				val: nums[r],
			})
		}
	}

	a := []int{0}
	Fprintln(out, nums[0])
	for len(a) != 0 {
		nxt := []int{}
		for _, v := range a {
			for _, vv := range g[v] {
				Fprintln(out, vv.val)
				nxt = append(nxt, vv.id)
			}
		}
		a = nxt
	}

}
