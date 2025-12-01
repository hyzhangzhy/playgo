package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var dfs func(rt *TreeNode) int
	dp := map[*TreeNode]int{}
	dfs = func(rt *TreeNode) int {
		if rt == nil {
			return 0
		}
		res, ok := dp[rt]
		if ok {
			return res
		}
		res1 := dfs(rt.Right) + dfs(rt.Left)
		res2 := rt.Val
		if rt.Right != nil {
			res2 += dfs(rt.Right.Left) + dfs(rt.Right.Right)
		}
		if rt.Left != nil {
			res2 += dfs(rt.Left.Left) + dfs(rt.Left.Right)
		}
		result := max(res1, res2)
		dp[rt] = result
		return result
	}
	return dfs(root)
}

func main() {

}
