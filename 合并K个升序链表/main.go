package main

import "container/heap"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

type heapItem struct {
	val  int
	node *ListNode
}

type nodeHeap []*heapItem

func (h nodeHeap) Len() int               { return len(h) }
func (h nodeHeap) Less(i, j int) bool     { return h[i].val < h[j].val }
func (h nodeHeap) Swap(i, j int)          { h[i], h[j] = h[j], h[i] }
func (h *nodeHeap) Push(item interface{}) { *h = append(*h, item.(*heapItem)) }
func (h *nodeHeap) Pop() interface{} {
	length := len(*h)
	result := (*h)[length-1]
	*h = (*h)[:length-1]
	return result
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 优先队列
	var nh nodeHeap
	for _, l := range lists {
		if l == nil {
			continue
		}
		nh = append(nh, &heapItem{val: l.Val, node: l})
	}
	heap.Init(&nh)

	dummyHead := &ListNode{}
	temp := dummyHead

	for len(nh) != 0 {
		minNode := heap.Pop(&nh).(*heapItem)
		temp.Next = minNode.node
		temp = temp.Next
		if minNode.node.Next != nil {
			heap.Push(&nh, &heapItem{val: minNode.node.Next.Val, node: minNode.node.Next})
		}
	}

	return dummyHead.Next
}

//func mergeKLists(lists []*ListNode) *ListNode {
//	newLists := make([]*ListNode, 0, len(lists))
//	for _, list := range lists {
//		newListNode := &ListNode{}
//		newListNode.Next = list
//		newLists = append(newLists, newListNode)
//	}
//
//	var resultHead *ListNode
//	var cur *ListNode
//
//	for {
//		var min *ListNode
//		for _, list := range newLists {
//			if list.Next == nil {
//				continue
//			}
//			if min == nil {
//				min = list
//			}
//			if min.Next.Val > list.Next.Val {
//				min = list
//			}
//		}
//
//		if min == nil {
//			break
//		}
//
//		if resultHead == nil {
//			resultHead = min.Next
//			cur = min.Next
//		} else {
//			cur.Next = min.Next
//			cur = cur.Next
//		}
//
//		min.Next = min.Next.Next
//	}
//
//	return resultHead
//}
