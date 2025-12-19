package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	res := &ListNode{}
	res.Next = head
	cur := head.Next
	// 通过nil把前面排好的切断
	head.Next = nil
	for cur != nil {
		nxt := cur.Next
		tmp := res.Next
		last := res
		for tmp != nil {
			if tmp.Val < cur.Val {
				last = tmp
				tmp = tmp.Next
			} else {
				break
			}
		}
		last.Next = cur
		cur.Next = tmp
		cur = nxt
	}

	return res.Next
}

func main() {

}
