package main

import "fmt"

func main() {

}

var (
	max = 0
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max = 0
	solve(root)
	return max
}

func solve(root *TreeNode) int {
	if root == nil {
		return -1
	}

	leftMax := solve(root.Left)
	rightMax := solve(root.Right)

	fmt.Println(leftMax, rightMax)

	if leftMax+rightMax+2 > max {
		max = leftMax+rightMax+2
	}

	return maxFunc(leftMax, rightMax) + 1
}

func maxFunc(i, j int) int {
	if i > j {
		return i
	}
	return j
}
