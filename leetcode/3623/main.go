package main

func countTrapezoids(points [][]int) int {
	//2
	//4
	//6
	//8
	//2*4 + 2*6 + 2*8
	//4*6 + 4*8
	//6*8
	mod := 1000000007
	mp := map[int]int{}
	for _, v := range points {
		mp[v[1]]++
	}
	values := []int{}
	for _, v := range mp {
		if v <= 1 {
			continue
		}
		test := v * (v - 1) / 2
		test %= mod
		values = append(values, test)
	}
	res := 0
	n := len(values)
	pSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		pSum[i] = pSum[i-1] + values[i-1]
		pSum[i] %= mod
	}
	for i := 0; i < n-1; i++ {
		add := values[i] * (pSum[n] - pSum[i+1]) % mod
		res += add
		res %= mod
	}
	return res
}

func main() {
	countTrapezoids([][]int{
		{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2},
	})
}
