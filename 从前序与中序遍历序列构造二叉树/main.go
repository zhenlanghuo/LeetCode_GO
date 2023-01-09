package main

import "sort"

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findRootIndex(inorder []int, root int) int {
	for i := 0; i < len(inorder); i++ {
		if inorder[i] == root {
			return i
		}
	}
	return 0
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{Val: preorder[0]}

	//rootIndex := findRootIndex(inorder, preorder[0])
	rootIndex := sort.SearchInts(inorder, preorder[0])

	root.Left = buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])

	return root
}
