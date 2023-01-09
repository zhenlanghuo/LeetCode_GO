package main

import "fmt"

func main() {
	list := []int{-1, 5, 3, 4, 0, 9, 3, 8, -5}
	dummyHead := &ListNode{}
	temp := dummyHead
	for _, val := range list {
		temp.Next = &ListNode{Val: val}
		temp = temp.Next
	}
	printList(dummyHead.Next)

	printList(sortList(dummyHead.Next))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func merge(list1, list2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	temp := dummyHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			temp.Next = list1
			list1 = list1.Next
		} else {
			temp.Next = list2
			list2 = list2.Next
		}
		temp = temp.Next
	}

	if list1 != nil {
		temp.Next = list1
	}

	if list2 != nil {
		temp.Next = list2
	}

	return dummyHead.Next
}

func sortList(head *ListNode) *ListNode {
	length := 0
	temp := head
	for temp != nil {
		length++
		temp = temp.Next
	}

	//fmt.Println(length)

	dummyHead := &ListNode{Next: head}
	for i := 1; i < length; i = i * 2 {
		cur, prev := dummyHead.Next, dummyHead
		for cur != nil {
			dummyHead1, dummyHead2 := &ListNode{}, &ListNode{}
			tempHead1, tempHead2 := dummyHead1, dummyHead2
			len1, len2 := 0, 0
			for j := 0; j < i; j++ {
				if cur == nil {
					break
				}
				tempHead1.Next = cur
				tempHead1 = tempHead1.Next
				cur = cur.Next
				len1++
			}
			tempHead1.Next = nil

			for j := 0; j < i; j++ {
				if cur == nil {
					break
				}
				tempHead2.Next = cur
				tempHead2 = tempHead2.Next
				cur = cur.Next
				len2++
			}
			tempHead2.Next = nil

			//printList(dummyHead1.Next)
			//printList(dummyHead2.Next)

			//fmt.Println(len1, len2)

			prev.Next = merge(dummyHead1.Next, dummyHead2.Next)
			//printList(prev.Next)
			//fmt.Println(prev.Val)
			for len1 > 0 || len2 > 0 {
				if len1 > 0 {
					len1 = len1 - 1
				} else if len2 > 0 {
					len2 = len2 - 1
				}
				prev = prev.Next
			}
			//fmt.Println(prev.Val)

		}
		//printList(dummyHead.Next)
	}

	return dummyHead.Next
}

func printList(head *ListNode) {
	list := make([]int, 0)
	temp := head
	for temp != nil {
		list = append(list, temp.Val)
		temp = temp.Next
	}
	fmt.Println(list)
}

//func sortList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil {
//		return head
//	}
//
//	fast, slow := head.Next, head
//	for fast != nil && fast.Next != nil {
//		slow = slow.Next
//		fast = fast.Next.Next
//	}
//
//	mid := slow.Next
//	slow.Next = nil
//
//	leftList := sortList(head)
//	rightList := sortList(mid)
//
//	return merge(leftList, rightList)
//}
