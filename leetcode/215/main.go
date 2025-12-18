package main

import (
	"fmt"
	"math/rand"
)

// 从大到小
// hoare分区
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	var qsel func(nums []int, l, r int, k int) int

	partition := func(nums []int, l, r int) int {
		//2 4 0 -3 2
		//rsel := rand.Int()%(r-l) + l
		partition := nums[r]
		i := l - 1
		j := r + 1
		for i < j {
			for i++; nums[i] > partition; i++ {
			}
			for j--; nums[j] < partition; j-- {
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		// <= partion 的最大下标
		return i - 1
	}

	//fmt.Println(partition([]int{2, -3, 0, 4, 2}, 0, 4))

	qsel = func(nums []int, l, r int, k int) int {
		if l == r {
			return nums[l]
		}
		rsel := rand.Int()%(r-l) + l
		nums[r], nums[rsel] = nums[rsel], nums[r]
		idx := partition(nums, l, r)
		if k <= idx {
			return qsel(nums, l, idx, k)
		} else {
			return qsel(nums, idx+1, r, k)
		}
	}

	return qsel(nums, 0, n-1, k-1)
}

func main() {
	fmt.Println(findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4))
}

/*
从小到大 n - k
func findKthLargest(nums []int, k int) int {
	n := len(nums)
	var qsel func(nums []int, l, r int, k int) int

	partition := func(nums []int, l, r int) int {
		//2 4 0 -3 2
		rsel := rand.Int()%(r-l) + l
		partition := nums[rsel]
		i := l - 1
		j := r + 1
		for i < j {
			for i++; nums[i] < partition; i++ {
			}
			for j--; nums[j] > partition; j-- {
			}
			if i < j {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
		// <= partion 的最大下标
		return i - 1
	}

	//fmt.Println(partition([]int{2, -3, 0, 4, 2}, 0, 4))

	qsel = func(nums []int, l, r int, k int) int {
		if l == r {
			return nums[l]
		}
		// rsel := rand.Int()%(r-l) + l
		// nums[r], nums[rsel] = nums[rsel], nums[r]
		idx := partition(nums, l, r)
		if k <= idx {
			return qsel(nums, l, idx, k)
		} else {
			return qsel(nums, idx+1, r, k)
		}
	}

	return qsel(nums, 0, n-1, n-k)
}
*/
