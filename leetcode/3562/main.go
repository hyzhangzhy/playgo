package main

import "fmt"

type dpstate struct {
	dp0 []int
	dp1 []int
}

func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	g := make([][]int, n+1)
	for _, v := range hierarchy {
		g[v[0]] = append(g[v[0]], v[1])
	}

	var dfs func(cur int) dpstate
	dfs = func(cur int) dpstate {
		nxtMax0 := make([]int, budget+1)
		nxtMax1 := make([]int, budget+1)
		// 分组背包问题求出他的所有的孩子如果当前选1，如果当前不选的最好结果
		for _, nxt := range g[cur] {
			son := dfs(nxt)
			for j := budget; j >= 0; j-- {
				for k := 0; k <= j; k++ {
					nxtMax0[j] = max(nxtMax0[j], son.dp0[k]+nxtMax0[j-k])
					nxtMax1[j] = max(nxtMax1[j], son.dp1[k]+nxtMax1[j-k])
				}
			}
		}
		cost := present[cur-1]
		get := future[cur-1]
		halfcost := cost / 2
		// cur 的父节点没有购买
		dp0 := make([]int, budget+1)
		dp1 := make([]int, budget+1)
		for i := 0; i <= budget; i++ {
			dp0[i] = nxtMax0[i]
			if i >= cost {
				dp0[i] = max(dp0[i], get-cost+nxtMax1[i-cost])
			}
			dp1[i] = nxtMax0[i]
			if i >= halfcost {
				dp1[i] = max(dp1[i], get-halfcost+nxtMax1[i-halfcost])
			}
		}

		return dpstate{dp0: dp0, dp1: dp1}
	}

	return dfs(1).dp0[budget]

}

func main() {
	fmt.Println(maxProfit(3, []int{5, 2, 3}, []int{8, 5, 6}, [][]int{{1, 2}, {2, 3}}, 7))
}
