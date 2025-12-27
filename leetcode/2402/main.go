package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type M struct {
	ed  int
	idx int
}

type Heap struct {
	data []M
}

func (h Heap) Len() int { return len(h.data) }
func (h Heap) Less(i, j int) bool {
	if h.data[i].ed == h.data[j].ed {
		return h.data[i].idx < h.data[j].idx
	}
	return h.data[i].ed < h.data[j].ed
}
func (h Heap) Swap(i, j int)       { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *Heap) Push(x interface{}) { h.data = append(h.data, x.(M)) }
func (h *Heap) Pop() interface{} {
	n := len(h.data)
	x := h.data[n-1]
	h.data = h.data[:n-1]
	return x
}
func (h *Heap) Top() M {
	return h.data[0]
}

func mostBooked(n int, meetings [][]int) int {
	res := make([]int, n)
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})
	nm := len(meetings)
	hp := &Heap{}
	isFree := make([]bool, n)
	for i := 0; i < n; i++ {
		isFree[i] = true
	}
	delay := 0
	for i := 0; i < nm; i++ {
		st := meetings[i][0]
		ed := meetings[i][1]
		for hp.Len() != 0 {
			top := hp.Top()
			if top.ed <= st {
				isFree[top.idx] = true
				heap.Pop(hp)
			} else {
				break
			}
		}
		if hp.Len() == n {
			top := hp.Top()
			delay = (top.ed - st)
			isFree[top.idx] = true
			heap.Pop(hp)
		}
		for j := 0; j < n; j++ {
			if isFree[j] {
				res[j]++
				heap.Push(hp, M{
					ed:  ed + delay,
					idx: j,
				})
				isFree[j] = false
				break
			}
		}
		delay = 0
	}
	idx := 0
	for i := 1; i < n; i++ {
		if res[i] > res[idx] {
			idx = i
		}
	}
	return idx
}

func main() {
	fmt.Println(mostBooked(2, [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}}))
}
