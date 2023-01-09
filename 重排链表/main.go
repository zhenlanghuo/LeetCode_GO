package main

type ListNode struct {
	Val  int
	Next *ListNode
}

//func reorderList(head *ListNode) {
//	if head == nil || head.Next == nil {
//		return
//	}
//
//	list := make([]*ListNode, 0)
//	temp := head
//	for temp != nil {
//		list = append(list, temp)
//		temp = temp.Next
//	}
//
//	i, j := 0, len(list)-1
//	prev := &ListNode{Next: head}
//	for i <= j {
//		if i != j {
//			list[i].Next = list[j]
//		}
//		prev.Next = list[i]
//		prev = list[j]
//		prev.Next = nil
//		i++
//		j--
//	}
//}

func reorderList(head *ListNode) {
	mid := middleNode(head)

	left := head
	right := mid.Next
	mid.Next = nil

	right = reserveNode(right)

	for right != nil {
		rightNext := right.Next
		leftNext := left.Next

		left.Next = right
		right.Next = leftNext


		left = leftNext
		right = rightNext
	}
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		slow = slow.Next
	}
	return slow
}

func reserveNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	next := head.Next
	head.Next = nil

	for next != nil {
		temp := next
		next = next.Next
		temp.Next = head
		head = temp
	}

	return head
}
