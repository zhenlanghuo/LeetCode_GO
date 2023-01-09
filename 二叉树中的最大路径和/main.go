package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	max := root.Val
	maxGain(root, &max)
	return max
}

func maxGain(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}

	leftGain := maxFunc(maxGain(root.Left, max), 0)
	rightGain := maxFunc(maxGain(root.Right, max), 0)

	*max = maxFunc(leftGain+rightGain+root.Val, *max)

	return maxFunc(leftGain+root.Val, rightGain+root.Val)
}

func maxFunc(x int, y int) int {
	if x > y {
		return x
	}
	return y
}
