package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func rob(root *TreeNode) int {
	// dp[x][0]表示x不可以偷, dp[x][1]表示x可以偷可以不偷
	dp := make(map[*TreeNode][]int)

	var dfs func(node *TreeNode, isCanSteal int) int

	dfs = func(node *TreeNode, isCanSteal int) int {
		if node == nil {
			return 0
		}

		if val, ok := dp[node]; ok {
			if val[isCanSteal] != -1 {
				return val[isCanSteal]
			}
		}

		maxn := math.MinInt32
		if isCanSteal == 1 {
			maxn = max(maxn, node.Val+dfs(node.Left, 0)+dfs(node.Right, 0))
		}
		maxn = max(maxn, dfs(node.Left, 1)+dfs(node.Right, 1))

		_, ok := dp[node]
		if !ok {
			dp[node] = []int{-1, -1}
		}
		dp[node][isCanSteal] = maxn

		return maxn
	}

	dfs(root, 1)

	return dp[root][1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
