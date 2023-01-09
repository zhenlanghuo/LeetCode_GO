package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

var sum int

func bstToGst(root *TreeNode) *TreeNode {
	sum = 0
	bstToGst_(root)

	return root

}

func bstToGst_(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	bstToGst_(root.Right)

	root.Val += sum
	sum = root.Val

	bstToGst_(root.Left)

	return root

}


func backtrack(root *TreeNode, n int) int {
	if root == nil {
		return n
	}

	//if root.Left == nil && root.Right == nil {
	//	root.Val += n
	//	return root.Val
	//}

	max := backtrack(root.Right, n)
	root.Val += max

	return backtrack(root.Left, root.Val)
}
