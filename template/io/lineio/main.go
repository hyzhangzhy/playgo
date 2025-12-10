package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
)

func lineIO() {
	in := bufio.NewScanner(os.Stdin) // 默认 4KB 初始 buffer
	in.Buffer(nil, math.MaxInt)      // 若单个 token 大小超过 65536 则加上这行，否则会报错
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for in.Scan() {
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

		fmt.Fprintln(out, nums)
		out.Flush()
		_ = sp
	}
}

func main() {
	lineIO()
}
