package main

func maxRotateFunction(nums []int) int {
	n := len(nums)
	sum := 0
	for i := 0; i < n; i++ {
		sum += nums[i]
	}
	maxm := 0
	zeroIdx := 0
	for i := 0; i < n; i++ {
		maxm += (i * nums[i])
	}
	// 25 + 2 * (9) - 6 * 3
	cur := maxm
	//fmt.Println(cur)
	for i := 1; i <= n-1; i++ {
		zeroIdx = (zeroIdx - 1 + n) % n
		nxt := cur + (sum - nums[zeroIdx]) - nums[zeroIdx]*(n-1)
		//fmt.Println(zeroIdx, sum- nums[zeroIdx], cur, nxt)
		maxm = max(maxm, nxt)
		cur = nxt
	}
	return maxm

	// 4 3 2 6
	// 0 1 2 3 0 1 2 3

	// 4 3 2 6
	// 1 2 3 0

	// 4 3 2 6
	// 2 3 0 1

	// 4 3 2 6
	// 3 0 1 2

	return maxm
}

func main() {

}
