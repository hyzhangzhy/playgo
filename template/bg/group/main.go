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

	var n, m int
	Fscan(in, &n, &m)

	// 定义物品结构体
	type Item struct {
		v, w int // v: 体积/花费, w: 价值/权重
	}

	// 使用切片存储分组，groups[i] 表示第 i 组的所有物品
	groups := make([][]Item, n)

	for i := 0; i < n; i++ {
		var s int
		Fscan(in, &s)
		// 为当前组分配切片空间
		groups[i] = make([]Item, s)
		for j := 0; j < s; j++ {
			Fscan(in, &groups[i][j].v, &groups[i][j].w)
		}
	}

	// dp 数组，f[j] 表示容量为 j 时的最大价值
	f := make([]int, m+1)

	// 分组背包逻辑
	for i := 0; i < n; i++ { // 枚举组
		for j := m; j >= 0; j-- { // 枚举容量（从大到小）
			for _, item := range groups[i] { // 枚举组内物品
				if item.v <= j {
					if val := f[j-item.v] + item.w; val > f[j] {
						f[j] = val
					}
				}
			}
		}
	}

	Fprintln(out, f[m])
}
