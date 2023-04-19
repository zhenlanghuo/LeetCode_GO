package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	ans := 0

	var dfs func(root *TreeNode) (int, int)
	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return math.MinInt64, math.MaxInt64
		}
		if root.Left == nil && root.Right == nil {
			return root.Val, root.Val
		}
		leftMx, leftMin := dfs(root.Left)
		rightMx, rightMin := dfs(root.Right)
		ans = max(ans, max(abs(root.Val-max(leftMx, rightMx)), abs(root.Val-min(leftMin, rightMin))))
		return max(root.Val, max(leftMx, rightMx)), min(root.Val, min(leftMin, rightMin))
	}
	dfs(root)

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
