package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	var root, node *ListNode

	add := 0
	num1 := 0
	num2 := 0

	for l1 != nil || l2 != nil {
		num1, num2 = 0, 0
		if l1 != nil {
			num1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			num2 = l2.Val
			l2 = l2.Next
		}

		sum := num1 + num2 + add
		if root == nil {
			root = &ListNode{Val: sum % 10}
			node = root
		} else {
			node.Next = &ListNode{Val: sum % 10}
			node = node.Next
		}

		add = sum / 10
	}

	if add != 0 {
		node.Next = &ListNode{Val: add}
	}


	return root
}
