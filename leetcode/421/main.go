package main

import (
	"fmt"
	"math/bits"
)

type Node struct {
	val int
	nxt [2]*Node
}

func findMaximumXOR(nums []int) int {
	//00101
	//11001
	//00010
	//01前缀树
	maxm := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		maxm = max(maxm, nums[i])
	}
	L := bits.Len(uint(maxm)) - 1
	root := &Node{}
	for i := 0; i < n; i++ {
		tmp := nums[i]
		curNode := root
		for j := L; j >= 0; j-- {
			cur := tmp & (1 << j)
			cur >>= j
			if curNode.nxt[cur] == nil {
				newNode := &Node{
					val: cur,
				}
				curNode.nxt[cur] = newNode
				curNode = newNode
			} else {
				curNode = curNode.nxt[cur]
			}
		}
	}
	res := 0
	for _, v := range nums {
		tmp := 0
		curNode := root
		for j := L; j >= 0; j-- {
			cur := v & (1 << j)
			cur >>= j
			if cur == 1 {
				if curNode.nxt[0] != nil {
					tmp |= (1 << j)
					curNode = curNode.nxt[0]
				} else {
					curNode = curNode.nxt[1]
				}
			} else {
				if curNode.nxt[1] != nil {
					tmp |= (1 << j)
					curNode = curNode.nxt[1]
				} else {
					curNode = curNode.nxt[0]
				}
			}
		}
		res = max(res, tmp)
	}

	return res

}

func main() {
	fmt.Println(findMaximumXOR([]int{3, 10, 5, 25, 2, 8}))
}
