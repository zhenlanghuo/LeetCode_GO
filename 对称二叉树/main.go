package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {

}

func inorder(root *TreeNode) string {
	if root == nil {
		return ""
	}

	leftStr := inorder(root.Left)

	rightStr := inorder(root.Right)

	return leftStr + strconv.FormatUint(uint64(root.Val), 10) + rightStr
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetric_(root.Left, root.Right)

}

func isSymmetric_(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	if !isSymmetric_(left.Left, right.Right) {
		return false
	}

	if !isSymmetric_(left.Right, right.Left) {
		return false
	}

	return true
}
