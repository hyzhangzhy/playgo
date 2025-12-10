package main

// 简单模板，只实现了路径压缩，没实现按秩合并。

func main() {
	n := 1000010

	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}

	find := func(x int) int {
		rt := x
		for fa[rt] != rt {
			rt = fa[rt]
		}
		for fa[x] != rt {
			fa[x], x = rt, fa[x]
		}
		return rt
	}

	unite := func(from, to int) {
		rt1 := find(from)
		rt2 := find(to)
		if rt1 == rt2 {
			return
		}
		fa[find(from)] = find(to)
	}

	unite(1, 4)

}
