package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	return solve(1, n)
}

func solve(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	result := make([]*TreeNode, 0)
	for i := start; i <= end; i++ {
		leftNodes := solve(start, i-1)
		rightNodes := solve(i+1, end)
		for _, left := range leftNodes {
			for _, right := range rightNodes {
				result = append(result, &TreeNode{Val: i, Left: left, Right: right})
			}
		}
	}

	return result
}
