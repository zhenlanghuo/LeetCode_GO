package main

import "fmt"

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{
				Val:   3,
			},
			Right: &TreeNode{
				Val:   4,
			},
		},
		Right: &TreeNode{
			Val: 5,
			Right: &TreeNode{
				Val: 6,
			},
		},
	}

	flatten(root)

	for root != nil {
		fmt.Println(root.Val)
		root = root.Right
	}

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//func flatten(root *TreeNode) {
//	flatten_(root)
//}

// 返回展开后最后的节点
//func flatten_(root *TreeNode) *TreeNode {
//	if root == nil {
//		return nil
//	}
//
//	left := root.Left
//	right := root.Right
//
//	// 左节点固定为nil
//	root.Left = nil
//
//	// lastNodeLeft 左节点展开后最后的节点
//	lastNodeLeft := flatten_(left)
//	if left != nil {
//		root.Right = left
//	}
//
//	// lastNodeLeft 右节点展开后最后的节点
//	lastNodeRight := flatten_(right)
//	if lastNodeLeft != nil {
//		lastNodeLeft.Right = right
//	}
//
//	// 如果右节点展开后最后的节点不为nil, 返回该节点
//	if lastNodeRight != nil {
//		return lastNodeRight
//	}
//	// 如果左节点展开后最后的节点不为nil, 返回该节点
//	if lastNodeLeft != nil {
//		return lastNodeLeft
//	}
//
//	// 如果左节点和右节点都为nil, 返回本身
//	return root
//}

func flatten(root *TreeNode) *TreeNode {
	cur := root
	for cur != nil {
		if cur.Left == nil {
			cur = cur.Right
			continue
		}

		predecessor := cur.Left
		for predecessor.Right != nil {
			predecessor = predecessor.Right
		}
		predecessor.Right = cur.Right
		cur.Right = cur.Left
		cur.Left = nil
	}
	return root
}
