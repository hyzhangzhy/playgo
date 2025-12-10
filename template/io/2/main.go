package main

import (
	"bufio"
	. "fmt"
	"os"
)

type ListNode struct {
	val  int
	next *ListNode
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	//res := 0
	head := &ListNode{}
	cur := head
	for i := 0; i < n; i++ {
		tmp := 0
		Fscan(in, &tmp)
		cur.next = &ListNode{val: tmp}
		cur = cur.next
	}
	cur = head.next
	for cur != nil {
		Fprintf(out, "%d\n", cur.val)
		cur = cur.next
	}
	//Fprintf(out, "%d", res)

}
