package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyHead := &ListNode{Next: head}
	temp := dummyHead
	count := 0
	var leftPrevNode, leftNode, rightNode, rightNextNode *ListNode
	for temp != nil {
		count++
		if count == left+1 {
			leftNode = temp
		}

		if count == left {
			leftPrevNode = temp
		}

		if count == right+1 {
			rightNode = temp
		}

		temp = temp.Next
	}
	rightNextNode = rightNode.Next
	rightNode.Next = nil

	reverseList(leftNode)

	leftNode.Next = rightNextNode
	leftPrevNode.Next = rightNode

	return dummyHead
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	if head.Next == nil {
		return head
	}
	next := head.Next
	root := reverseList(head.Next)
	next.Next = head
	head.Next = nil

	return root
}
