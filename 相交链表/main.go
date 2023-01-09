package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

//func getIntersectionNode(headA, headB *ListNode) *ListNode {
//	listNodeMap := make(map[*ListNode]bool)
//	for headA != nil {
//		listNodeMap[headA] = true
//		headA = headA.Next
//	}
//
//	for headB != nil {
//		if _, ok := listNodeMap[headB]; ok {
//			return headB
//		}
//		headB = headB.Next
//	}
//
//	return nil
//}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}

	return pa
}
