package main

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	fast, slow := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast != slow {
			continue
		}
		slow2 := head
		for slow != slow2 {
			slow = slow.Next
			slow2 = slow2.Next
		}
		return slow2
	}

	return nil
}
