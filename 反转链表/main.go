package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
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
