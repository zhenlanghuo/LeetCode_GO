package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == key {
		if root.Left == nil {
			return root.Right
		} else if root.Right == nil {
			return root.Left
		} else {
			left := root.Right
			for left.Left != nil {
				left = left.Left
			}
			left.Left = root.Left
			return root.Right


			right := root.Left
			for right.Right != nil {
				right = right.Right
			}
			right.Right = root.Right
			return root.Left
		}
	}

	root.Left = deleteNode(root.Left, key)
	root.Right = deleteNode(root.Right, key)

	return root
}
