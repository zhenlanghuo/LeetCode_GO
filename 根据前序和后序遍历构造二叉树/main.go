package main

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findRootIndex(order []int, root int) int {
	for i := 0; i < len(order); i++ {
		if order[i] == root {
			return i
		}
	}
	return 0
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {

	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	preorder = preorder[1:]
	postorder = postorder[:len(postorder)-1]

	if len(preorder) == 0 {
		return root
	}
	index := findRootIndex(postorder, preorder[0])

	root.Left = constructFromPrePost(preorder[:index+1], postorder[:index+1])
	root.Right = constructFromPrePost(preorder[index+1:], postorder[index+1:])

	return root
}
