package main

func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	// x + y + a + b = 0
	//x + y = -(a+b)
	mp1 := map[int]int{}
	n := len(nums1)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			test := nums1[i] + nums2[j]
			mp1[test]++
		}
	}
	mp2 := map[int]int{}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			test := nums3[i] + nums4[j]
			mp2[test]++
		}
	}
	res := 0
	for k, v := range mp1 {
		x := -k
		val, ok := mp2[x]
		if ok {
			res += v * val
		}
	}
	return res
}

func main() {
	fourSumCount([]int{1, 2}, []int{-2, -1}, []int{-1, 2}, []int{0, 2})
}
