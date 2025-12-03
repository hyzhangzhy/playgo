package main

import (
	"fmt"
	"math"
)

type P struct {
	x int
	y int
}

func countTrapezoids(points [][]int) int {
	// 斜率+截距分组，统计梯形个数，用枚举右维护左的思想，因为这样能O(1)的递推
	// 用相同的中点不同的斜率来构成平行四边形
	// 结果为梯形的个数减去平行四边形的个数
	n := len(points)
	// k -> b -> count 统计梯形
	trap := map[float64]map[float64]int{}
	// 中点 -> k -> count 统计平行四边形
	par := map[P]map[float64]int{}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			x1, y1 := points[i][0], points[i][1]
			x2, y2 := points[j][0], points[j][1]
			dy := y2 - y1
			dx := x2 - x1
			k := math.MaxFloat64
			b := float64(x1)
			if dx != 0 {
				fdy, fdx, fy, fx := float64(dy), float64(dx), float64(y1), float64(x1)
				k = fdy / fdx
				b = (fy*fdx - fx*fdy) / fdx
			}
			_, ok := trap[k]
			if !ok {
				trap[k] = map[float64]int{}
			}
			trap[k][b]++
			mid := P{
				x: x1 + x2,
				y: y1 + y2,
			}
			_, ok = par[mid]
			if !ok {
				par[mid] = map[float64]int{}
			}
			par[mid][k]++
		}
	}
	sum1 := 0
	for _, v := range trap {
		left := 0
		for _, v1 := range v {
			sum1 += left * v1
			left += v1
		}
	}
	sum2 := 0
	for _, v := range par {
		left2 := 0
		for _, v1 := range v {
			sum2 += left2 * v1
			left2 += v1
		}
	}
	return sum1 - sum2
}

func main() {
	fmt.Println(countTrapezoids([][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}}))
}
