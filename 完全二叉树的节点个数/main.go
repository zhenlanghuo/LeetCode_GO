package main

import "math"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	lh, rh := 0, 0
	temp := root
	for temp != nil {
		lh++
		temp = temp.Left
	}
	temp = root
	for temp != nil {
		rh++
		temp = temp.Right
	}
	if lh == rh {
		return int(math.Pow(2, float64(lh)) - 1)
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}
