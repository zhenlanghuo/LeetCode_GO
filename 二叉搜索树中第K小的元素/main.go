package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var cur = 0
var result = 0

func kthSmallest(root *TreeNode, k int) int {
	cur = 0
	result = 0
	kthSmallest_(root, k)
	return result
}

func kthSmallest_(root *TreeNode, k int) {
	if root == nil {
		return
	}

	kthSmallest_(root.Left, k)
	cur++
	if cur == k {
		result = root.Val
	}
	kthSmallest_(root.Right, k)
}
