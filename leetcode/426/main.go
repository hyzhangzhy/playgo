package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	nodes := []*Node{}
	var dfs func(rt *Node)
	dfs = func(rt *Node) {
		if rt == nil {
			return
		}
		dfs(rt.Left)
		nodes = append(nodes, rt)
		dfs(rt.Right)
	}
	dfs(root)
	n := len(nodes)
	for i := 0; i < n; i++ {
		nodes[i].Right = nodes[(i+1)%n]
		nodes[i].Left = nodes[(i-1+n)%n]
	}
	return nodes[0]
}

func main() {

}
