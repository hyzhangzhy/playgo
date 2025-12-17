package main

import (
	"fmt"
	"sync"
)

type MP struct {
	mp   map[string]int
	lock sync.RWMutex
}

func New() *MP {
	return &MP{
		mp:   map[string]int{},
		lock: sync.RWMutex{},
	}
}

func (m *MP) Get(k string) (int, bool) {
	m.lock.RLock()
	res, ok := m.mp[k]
	m.lock.RUnlock()
	return res, ok
}

func (m *MP) Set(k string, v int) {
	m.lock.Lock()
	m.mp[k] = v
	m.lock.Unlock()
}

//tail -f xxx.log -n 1000 | grep error
//

//大的key删除导致阻塞

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	cur := head
	var prev *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}

func main() {
	test := []int{1, 2, 3}
	head := &ListNode{}
	cur := head
	for i := 0; i < 3; i++ {
		cur.Next = &ListNode{
			Val: test[i],
		}
		cur = cur.Next
	}
	x := reverseList(head.Next)
	for x != nil {
		fmt.Println(x.Val)
		x = x.Next
	}

}
