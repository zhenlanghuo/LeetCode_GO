package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	smallHead, largeHead := &ListNode{Next: head}, &ListNode{Next: head}
	smallTail, largeTail := smallHead, largeHead

	temp := head
	for temp != nil {
		if temp.Val < x {
			smallTail.Next = temp
			smallTail = smallTail.Next
		} else {
			largeTail.Next = temp
			largeTail = largeTail.Next
		}

		temp = temp.Next
	}

	largeTail.Next = nil
	smallTail.Next = largeHead

	return smallHead.Next
}

//func partition(head *ListNode, x int) *ListNode {
//
//	dummyHead := &ListNode{Next: head}
//
//	lessTail := dummyHead
//	for lessTail != nil && lessTail.Next != nil && lessTail.Next.Val < x {
//		lessTail = lessTail.Next
//	}
//
//	if lessTail == nil {
//		return lessTail
//	}
//
//	moreTail := lessTail.Next
//	for moreTail != nil {
//		for moreTail != nil && moreTail.Next != nil && moreTail.Next.Val >= x {
//			moreTail = moreTail.Next
//		}
//
//		if moreTail == nil || moreTail.Next == nil {
//			break
//		}
//
//		tempHead := moreTail.Next
//		temp := tempHead
//		for temp != nil && temp.Next != nil && temp.Next.Val < x {
//			temp = temp.Next
//		}
//
//		if temp == nil {
//			break
//		}
//
//		moreTail.Next = temp.Next
//		moreTail = moreTail.Next
//
//		temp.Next = lessTail.Next
//		lessTail.Next = tempHead
//		lessTail = temp
//	}
//
//
//	return dummyHead.Next
//}
