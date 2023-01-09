// K 个一组翻转链表
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	dummyHead := &ListNode{Next: head}
	prev := dummyHead
	for head != nil {
		tail := head
		for i := 0; i < k-1; i++ {
			tail = tail.Next
			if tail == nil {
				return dummyHead.Next
			}
		}

		tempHead, tempTail := reverseMyGroup(head, tail)
		prev.Next = tempHead
		prev = tempTail
		head = tempTail.Next
	}

	return dummyHead.Next
}

func reverseMyGroup(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	p := head
	for prev != tail {
		next := p.Next
		p.Next = prev
		prev = p
		p = next
	}
	return tail, head
}
